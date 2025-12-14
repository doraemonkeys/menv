package env

import (
	"os/exec"
	"sort"
	"strings"
)

const (
	userEnvRegPath   = "HKEY_CURRENT_USER\\Environment"
	systemEnvRegPath = "HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Session Manager\\Environment"
)

// EnvVar represents an environment variable with key and value.
type EnvVar struct {
	Key   string
	Value string
}

// ListUser lists all user environment variables from registry.
func ListUser() ([]EnvVar, error) {
	return listFromRegistry(userEnvRegPath)
}

// ListSystem lists all system environment variables from registry.
func ListSystem() ([]EnvVar, error) {
	return listFromRegistry(systemEnvRegPath)
}

// GetUser gets a specific user environment variable value.
func GetUser(key string) (string, error) {
	return getFromRegistry(userEnvRegPath, key)
}

// GetSystem gets a specific system environment variable value.
func GetSystem(key string) (string, error) {
	return getFromRegistry(systemEnvRegPath, key)
}

func listFromRegistry(regPath string) ([]EnvVar, error) {
	cmd := exec.Command("reg", "query", regPath)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return parseRegOutput(string(output)), nil
}

func getFromRegistry(regPath, key string) (string, error) {
	cmd := exec.Command("reg", "query", regPath, "/v", key)
	output, err := cmd.Output()
	if err != nil {
		// Key not found is not an error, just return empty
		return "", nil
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "HKEY_") {
			continue
		}
		parts := parseRegLine(line)
		if parts != nil && strings.EqualFold(parts.Key, key) {
			return parts.Value, nil
		}
	}
	return "", nil
}

func parseRegOutput(output string) []EnvVar {
	lines := strings.Split(output, "\n")
	var result []EnvVar

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "HKEY_") {
			continue
		}
		if env := parseRegLine(line); env != nil {
			result = append(result, *env)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return strings.ToLower(result[i].Key) < strings.ToLower(result[j].Key)
	})

	return result
}

// SearchUser searches user env vars by keyword (case-insensitive, matches key or value).
func SearchUser(keyword string) ([]EnvVar, error) {
	envVars, err := ListUser()
	if err != nil {
		return nil, err
	}
	return filterByKeyword(envVars, keyword), nil
}

// SearchSystem searches system env vars by keyword (case-insensitive, matches key or value).
func SearchSystem(keyword string) ([]EnvVar, error) {
	envVars, err := ListSystem()
	if err != nil {
		return nil, err
	}
	return filterByKeyword(envVars, keyword), nil
}

func filterByKeyword(envVars []EnvVar, keyword string) []EnvVar {
	keyword = strings.ToLower(keyword)
	var result []EnvVar
	for _, e := range envVars {
		if strings.Contains(strings.ToLower(e.Key), keyword) ||
			strings.Contains(strings.ToLower(e.Value), keyword) {
			result = append(result, e)
		}
	}
	return result
}

// parseRegLine parses a line like "    JAVA_HOME    REG_SZ    C:\Java\jdk"
func parseRegLine(line string) *EnvVar {
	// Registry output format: NAME    TYPE    VALUE
	// Types: REG_SZ, REG_EXPAND_SZ, etc.
	fields := strings.Fields(line)
	if len(fields) < 3 {
		return nil
	}

	key := fields[0]
	regType := fields[1]

	// Only handle string types
	if regType != "REG_SZ" && regType != "REG_EXPAND_SZ" {
		return nil
	}

	// Value may contain spaces, so rejoin everything after the type
	valueStart := strings.Index(line, regType)
	if valueStart == -1 {
		return nil
	}
	value := strings.TrimSpace(line[valueStart+len(regType):])

	return &EnvVar{Key: key, Value: value}
}
