package path

import (
	"errors"
	"os/exec"
	"strings"
)

const (
	userEnvRegPath   = "HKEY_CURRENT_USER\\Environment"
	systemEnvRegPath = "HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Session Manager\\Environment"
)

// QueryUserPath queries the user's PATH environment variable from registry.
func QueryUserPath() ([]string, error) {
	pathCmd := exec.Command("reg", "query", userEnvRegPath, "/v", "Path")
	pathByte, err := pathCmd.Output()
	if err != nil {
		return nil, err
	}
	return parsePathOutput(string(pathByte))
}

// QuerySystemPath queries the system PATH environment variable from registry.
func QuerySystemPath() ([]string, error) {
	pathCmd := exec.Command("reg", "query", systemEnvRegPath, "/v", "Path")
	pathByte, err := pathCmd.Output()
	if err != nil {
		return nil, err
	}
	return parsePathOutput(string(pathByte))
}

func parsePathOutput(output string) ([]string, error) {
	output = strings.TrimSpace(output)

	// Find Path line and extract value after REG_SZ or REG_EXPAND_SZ
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(strings.ToLower(line), "path") {
			continue
		}

		// Try REG_SZ first, then REG_EXPAND_SZ
		for _, regType := range []string{"REG_EXPAND_SZ", "REG_SZ"} {
			if idx := strings.Index(line, regType); idx != -1 {
				pathValue := strings.TrimSpace(line[idx+len(regType):])
				return splitAndCleanPath(pathValue), nil
			}
		}
	}

	return nil, errors.New("query path failed")
}

// splitAndCleanPath splits path by semicolon and removes empty entries.
func splitAndCleanPath(path string) []string {
	paths := strings.Split(path, ";")
	result := make([]string, 0, len(paths))

	for _, p := range paths {
		p = strings.TrimSpace(p)
		if p != "" {
			result = append(result, p)
		}
	}

	return result
}
