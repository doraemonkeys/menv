package env

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type ExportFormat string

const (
	FormatShell ExportFormat = "sh"
	FormatBatch ExportFormat = "bat"
	FormatJSON  ExportFormat = "json"
)

func DetectFormat(filename string) ExportFormat {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".sh":
		return FormatShell
	case ".bat", ".cmd":
		return FormatBatch
	case ".json":
		return FormatJSON
	default:
		return FormatShell
	}
}

func Export(filename string, envVars []EnvVar) error {
	format := DetectFormat(filename)

	var content string
	var err error

	switch format {
	case FormatShell:
		content = formatShell(envVars)
	case FormatBatch:
		content = formatBatch(envVars)
	case FormatJSON:
		content, err = formatJSON(envVars)
		if err != nil {
			return err
		}
	}

	return os.WriteFile(filename, []byte(content), 0644)
}

func formatShell(envVars []EnvVar) string {
	var sb strings.Builder
	sb.WriteString("#!/bin/bash\n\n")
	for _, e := range envVars {
		value := strings.ReplaceAll(e.Value, `"`, `\"`)
		sb.WriteString(fmt.Sprintf("export %s=\"%s\"\n", e.Key, value))
	}
	return sb.String()
}

func formatBatch(envVars []EnvVar) string {
	var sb strings.Builder
	sb.WriteString("@echo off\r\n\r\n")
	for _, e := range envVars {
		sb.WriteString(fmt.Sprintf("SET %s=%s\r\n", e.Key, e.Value))
	}
	return sb.String()
}

func formatJSON(envVars []EnvVar) (string, error) {
	envMap := make(map[string]string)
	for _, e := range envVars {
		envMap[e.Key] = e.Value
	}

	data, err := json.MarshalIndent(envMap, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data) + "\n", nil
}
