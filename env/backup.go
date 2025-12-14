package env

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type BackupData struct {
	CreatedAt time.Time `json:"created_at"`
	Source    string    `json:"source"`
	EnvVars   []EnvVar  `json:"env_vars"`
}

func Backup(filename string, isSystem bool) (int, error) {
	var envVars []EnvVar
	var err error
	var source string

	if isSystem {
		envVars, err = ListSystem()
		source = "system"
	} else {
		envVars, err = ListUser()
		source = "user"
	}

	if err != nil {
		return 0, err
	}

	backup := BackupData{
		CreatedAt: time.Now(),
		Source:    source,
		EnvVars:   envVars,
	}

	data, err := json.MarshalIndent(backup, "", "  ")
	if err != nil {
		return 0, err
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return 0, err
	}

	return len(envVars), nil
}

func Restore(filename string, isSystem bool) (int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	var backup BackupData
	if err := json.Unmarshal(data, &backup); err != nil {
		return 0, fmt.Errorf("invalid backup file: %w", err)
	}

	for _, e := range backup.EnvVars {
		if isSystem {
			if err := SetSystem(e.Key, e.Value); err != nil {
				return 0, fmt.Errorf("failed to set %s: %w", e.Key, err)
			}
		} else {
			if err := Set(e.Key, e.Value); err != nil {
				return 0, fmt.Errorf("failed to set %s: %w", e.Key, err)
			}
		}
	}

	return len(backup.EnvVars), nil
}

func LoadBackup(filename string) (*BackupData, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var backup BackupData
	if err := json.Unmarshal(data, &backup); err != nil {
		return nil, fmt.Errorf("invalid backup file: %w", err)
	}

	return &backup, nil
}
