package path

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// Add adds a path to the PATH environment variable.
// If sys is true, modifies system PATH; otherwise modifies user PATH.
// Note: setx has a 1024 character limit, so PowerShell is used instead.
func Add(add string, sys bool) error {
	if strings.ContainsRune(add, ';') {
		return errors.New("invalid path: " + add)
	}

	var (
		paths []string
		err   error
	)

	if sys {
		paths, err = QuerySystemPath()
	} else {
		paths, err = QueryUserPath()
	}
	if err != nil {
		return err
	}

	// Normalize the path to add
	add = normalizePath(add)

	// Check if path already exists
	var newPath string
	for _, p := range paths {
		p = normalizePath(p)
		if p == add {
			fmt.Printf("skip %s\n", add)
			return nil
		}
		newPath += p + ";"
	}

	fmt.Printf("add %s\n", add)
	newPath += add + ";"

	// Use PowerShell to set the path (avoids setx 1024 char limit)
	var cmd *exec.Cmd
	if sys {
		cmd = exec.Command("powershell", "[System.Environment]::SetEnvironmentVariable(\"Path\", \""+newPath+"\", \"Machine\")")
	} else {
		cmd = exec.Command("powershell", "[System.Environment]::SetEnvironmentVariable(\"Path\", \""+newPath+"\", \"User\")")
	}

	return cmd.Run()
}

// CleanUser removes duplicate entries from user PATH.
func CleanUser() error {
	pathCmd := exec.Command("reg", "query", userEnvRegPath, "/v", "Path")
	pathByte, err := pathCmd.Output()
	if err != nil {
		return err
	}

	path := strings.TrimSpace(string(pathByte))
	paths := strings.Split(path, ";")
	seen := make(map[string]bool, len(paths))
	var newPath string

	for _, p := range paths {
		p = strings.TrimSpace(p)
		if p != "" && !seen[p] {
			seen[p] = true
			newPath += p + ";"
			fmt.Println(p)
		}
	}

	cmd := exec.Command("powershell", "[System.Environment]::SetEnvironmentVariable(\"Path\", \""+newPath+"\", \"User\")")
	return cmd.Run()
}

// normalizePath removes trailing slashes and trims whitespace.
func normalizePath(p string) string {
	p = strings.TrimSpace(p)
	p = strings.TrimSuffix(p, "\\")
	p = strings.TrimSuffix(p, "/")
	return strings.TrimSpace(p)
}
