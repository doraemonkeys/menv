package env

import (
	"bytes"
	"errors"
	"strings"

	"github.com/doraemonkeys/doraemon"
)

// ParseEnvFile parses environment file content and returns key-value pairs.
// It supports optional prefix filtering with startWith parameter.
func ParseEnvFile(content []byte, startWith string) ([]doraemon.Pair[string, string], error) {
	envMap := []doraemon.Pair[string, string]{}
	lines := bytes.Split(content, []byte("\n"))

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(string(line))

		// Ignore empty lines and comments
		if len(trimmedLine) == 0 || strings.HasPrefix(trimmedLine, "#") {
			continue
		}

		// Filter by prefix if specified
		if startWith != "" && !strings.HasPrefix(trimmedLine, startWith) {
			continue
		}

		// Remove prefix (e.g., "export ")
		if startWith != "" {
			trimmedLine = strings.TrimPrefix(trimmedLine, startWith+" ")
		}

		// Split key and value by '='
		parts := strings.SplitN(trimmedLine, "=", 2)
		if len(parts) != 2 {
			return nil, errors.New("invalid format: missing '=' in line: " + trimmedLine)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if strings.Contains(value, "=") {
			return nil, errors.New("invalid format: multiple '=' in line: " + trimmedLine)
		}

		envMap = append(envMap, doraemon.Pair[string, string]{First: key, Second: value})
	}

	return envMap, nil
}
