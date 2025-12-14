package path

import (
	"errors"
	"os/exec"
	"strings"
)

const (
	userEnvRegPath   = "HKEY_CURRENT_USER\\Environment"
	systemEnvRegPath = "HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Session Manager\\Environment"
	pathKeyword      = "Path    REG_SZ"
)

// QueryUserPath queries the user's PATH environment variable from registry.
func QueryUserPath() ([]string, error) {
	pathCmd := exec.Command("reg", "query", userEnvRegPath, "/v", "Path")
	pathByte, err := pathCmd.Output()
	if err != nil {
		return nil, err
	}

	nowPath := strings.TrimSpace(string(pathByte))
	if !strings.Contains(nowPath, pathKeyword) {
		return nil, errors.New("query path failed, get path: " + nowPath)
	}

	nowPath = strings.TrimSpace(nowPath[strings.LastIndex(nowPath, pathKeyword)+len(pathKeyword):])
	return splitAndCleanPath(nowPath), nil
}

// QuerySystemPath queries the system PATH environment variable from registry.
func QuerySystemPath() ([]string, error) {
	pathCmd := exec.Command("reg", "query", systemEnvRegPath, "/v", "Path")
	pathByte, err := pathCmd.Output()
	if err != nil {
		return nil, err
	}

	nowPath := strings.TrimSpace(string(pathByte))
	if !strings.Contains(nowPath, pathKeyword) {
		return nil, errors.New("query path failed")
	}

	nowPath = strings.TrimSpace(nowPath[strings.LastIndex(nowPath, pathKeyword)+len(pathKeyword):])
	return splitAndCleanPath(nowPath), nil
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
