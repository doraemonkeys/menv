package env

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDetectFormat(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     ExportFormat
	}{
		{
			name:     "shell script .sh",
			filename: "env.sh",
			want:     FormatShell,
		},
		{
			name:     "shell script uppercase",
			filename: "ENV.SH",
			want:     FormatShell,
		},
		{
			name:     "batch file .bat",
			filename: "env.bat",
			want:     FormatBatch,
		},
		{
			name:     "batch file .cmd",
			filename: "env.cmd",
			want:     FormatBatch,
		},
		{
			name:     "json file",
			filename: "env.json",
			want:     FormatJSON,
		},
		{
			name:     "unknown extension defaults to shell",
			filename: "env.txt",
			want:     FormatShell,
		},
		{
			name:     "no extension defaults to shell",
			filename: "envfile",
			want:     FormatShell,
		},
		{
			name:     "path with directory",
			filename: "/some/path/env.bat",
			want:     FormatBatch,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DetectFormat(tt.filename)
			if got != tt.want {
				t.Errorf("DetectFormat(%q) = %v, want %v", tt.filename, got, tt.want)
			}
		})
	}
}

func TestExport(t *testing.T) {
	envVars := []EnvVar{
		{Key: "FOO", Value: "bar"},
		{Key: "BAZ", Value: "qux"},
	}

	tests := []struct {
		name       string
		filename   string
		envVars    []EnvVar
		wantPrefix string
	}{
		{
			name:       "export to shell",
			filename:   "test.sh",
			envVars:    envVars,
			wantPrefix: "#!/bin/bash",
		},
		{
			name:       "export to batch",
			filename:   "test.bat",
			envVars:    envVars,
			wantPrefix: "@echo off",
		},
		{
			name:       "export to json",
			filename:   "test.json",
			envVars:    envVars,
			wantPrefix: "{",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			filePath := filepath.Join(tmpDir, tt.filename)

			err := Export(filePath, tt.envVars)
			if err != nil {
				t.Fatalf("Export() error = %v", err)
			}

			content, err := os.ReadFile(filePath)
			if err != nil {
				t.Fatalf("Failed to read exported file: %v", err)
			}

			if !strings.HasPrefix(string(content), tt.wantPrefix) {
				t.Errorf("Export() content prefix = %q, want prefix %q", string(content)[:20], tt.wantPrefix)
			}
		})
	}
}

func TestFormatShell(t *testing.T) {
	tests := []struct {
		name    string
		envVars []EnvVar
		want    string
	}{
		{
			name:    "empty vars",
			envVars: []EnvVar{},
			want:    "#!/bin/bash\n\n",
		},
		{
			name:    "single var",
			envVars: []EnvVar{{Key: "FOO", Value: "bar"}},
			want:    "#!/bin/bash\n\nexport FOO=\"bar\"\n",
		},
		{
			name:    "multiple vars",
			envVars: []EnvVar{{Key: "FOO", Value: "bar"}, {Key: "BAZ", Value: "qux"}},
			want:    "#!/bin/bash\n\nexport FOO=\"bar\"\nexport BAZ=\"qux\"\n",
		},
		{
			name:    "value with quotes",
			envVars: []EnvVar{{Key: "MSG", Value: `say "hello"`}},
			want:    "#!/bin/bash\n\nexport MSG=\"say \\\"hello\\\"\"\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatShell(tt.envVars)
			if got != tt.want {
				t.Errorf("formatShell() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestFormatBatch(t *testing.T) {
	tests := []struct {
		name    string
		envVars []EnvVar
		want    string
	}{
		{
			name:    "empty vars",
			envVars: []EnvVar{},
			want:    "@echo off\r\n\r\n",
		},
		{
			name:    "single var",
			envVars: []EnvVar{{Key: "FOO", Value: "bar"}},
			want:    "@echo off\r\n\r\nSET FOO=bar\r\n",
		},
		{
			name:    "multiple vars",
			envVars: []EnvVar{{Key: "FOO", Value: "bar"}, {Key: "BAZ", Value: "qux"}},
			want:    "@echo off\r\n\r\nSET FOO=bar\r\nSET BAZ=qux\r\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatBatch(tt.envVars)
			if got != tt.want {
				t.Errorf("formatBatch() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestFormatJSON(t *testing.T) {
	tests := []struct {
		name    string
		envVars []EnvVar
		wantErr bool
	}{
		{
			name:    "empty vars",
			envVars: []EnvVar{},
			wantErr: false,
		},
		{
			name:    "single var",
			envVars: []EnvVar{{Key: "FOO", Value: "bar"}},
			wantErr: false,
		},
		{
			name:    "multiple vars",
			envVars: []EnvVar{{Key: "FOO", Value: "bar"}, {Key: "BAZ", Value: "qux"}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := formatJSON(tt.envVars)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if !strings.HasPrefix(got, "{") || !strings.HasSuffix(got, "}\n") {
					t.Errorf("formatJSON() = %q, want valid JSON", got)
				}
			}
		})
	}
}

func TestFormatJSONContent(t *testing.T) {
	envVars := []EnvVar{{Key: "FOO", Value: "bar"}}
	got, err := formatJSON(envVars)
	if err != nil {
		t.Fatalf("formatJSON() error = %v", err)
	}

	if !strings.Contains(got, `"FOO"`) || !strings.Contains(got, `"bar"`) {
		t.Errorf("formatJSON() = %q, want to contain FOO and bar", got)
	}
}
