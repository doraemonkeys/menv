package env

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestBackupData_Marshal(t *testing.T) {
	backup := BackupData{
		CreatedAt: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
		Source:    "user",
		EnvVars:   []EnvVar{{Key: "FOO", Value: "bar"}, {Key: "BAZ", Value: "qux"}},
	}

	data, err := json.MarshalIndent(backup, "", "  ")
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}

	var restored BackupData
	if err := json.Unmarshal(data, &restored); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}

	if restored.Source != backup.Source {
		t.Errorf("Source = %q, want %q", restored.Source, backup.Source)
	}

	if len(restored.EnvVars) != len(backup.EnvVars) {
		t.Errorf("EnvVars count = %d, want %d", len(restored.EnvVars), len(backup.EnvVars))
	}

	for i, e := range restored.EnvVars {
		if e.Key != backup.EnvVars[i].Key || e.Value != backup.EnvVars[i].Value {
			t.Errorf("EnvVar[%d] = {%q, %q}, want {%q, %q}",
				i, e.Key, e.Value, backup.EnvVars[i].Key, backup.EnvVars[i].Value)
		}
	}
}

func TestLoadBackup(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name        string
		content     string
		wantErr     bool
		wantSource  string
		wantEnvVars int
	}{
		{
			name: "valid backup",
			content: `{
				"created_at": "2024-01-01T12:00:00Z",
				"source": "user",
				"env_vars": [
					{"Key": "FOO", "Value": "bar"},
					{"Key": "BAZ", "Value": "qux"}
				]
			}`,
			wantErr:     false,
			wantSource:  "user",
			wantEnvVars: 2,
		},
		{
			name: "system backup",
			content: `{
				"created_at": "2024-01-01T12:00:00Z",
				"source": "system",
				"env_vars": [
					{"Key": "PATH", "Value": "C:\\Windows"}
				]
			}`,
			wantErr:     false,
			wantSource:  "system",
			wantEnvVars: 1,
		},
		{
			name: "empty env vars",
			content: `{
				"created_at": "2024-01-01T12:00:00Z",
				"source": "user",
				"env_vars": []
			}`,
			wantErr:     false,
			wantSource:  "user",
			wantEnvVars: 0,
		},
		{
			name:    "invalid json",
			content: `{invalid}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := filepath.Join(tmpDir, tt.name+".json")
			if err := os.WriteFile(filePath, []byte(tt.content), 0644); err != nil {
				t.Fatalf("Failed to write test file: %v", err)
			}

			backup, err := LoadBackup(filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadBackup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if backup.Source != tt.wantSource {
				t.Errorf("Source = %q, want %q", backup.Source, tt.wantSource)
			}

			if len(backup.EnvVars) != tt.wantEnvVars {
				t.Errorf("EnvVars count = %d, want %d", len(backup.EnvVars), tt.wantEnvVars)
			}
		})
	}
}

func TestLoadBackup_FileNotFound(t *testing.T) {
	_, err := LoadBackup("/nonexistent/path/backup.json")
	if err == nil {
		t.Error("LoadBackup() expected error for nonexistent file")
	}
}
