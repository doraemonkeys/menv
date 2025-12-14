package path

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/doraemonkeys/menv/color"
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
			color.Warning("skip %s (already exists)", add)
			return nil
		}
		newPath += p + ";"
	}

	newPath += add + ";"

	// Use PowerShell to set the path (avoids setx 1024 char limit)
	var cmd *exec.Cmd
	scope := "user"
	if sys {
		scope = "system"
		cmd = exec.Command("powershell", "[System.Environment]::SetEnvironmentVariable(\"Path\", \""+newPath+"\", \"Machine\")")
	} else {
		cmd = exec.Command("powershell", "[System.Environment]::SetEnvironmentVariable(\"Path\", \""+newPath+"\", \"User\")")
	}

	if err := cmd.Run(); err != nil {
		return err
	}
	color.Success("add  %s [%s PATH]", add, scope)
	return nil
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
	removed := 0

	for _, p := range paths {
		p = strings.TrimSpace(p)
		if p != "" && !seen[p] {
			seen[p] = true
			newPath += p + ";"
		} else if p != "" {
			removed++
		}
	}

	cmd := exec.Command("powershell", "[System.Environment]::SetEnvironmentVariable(\"Path\", \""+newPath+"\", \"User\")")
	if err := cmd.Run(); err != nil {
		return err
	}
	color.Success("cleaned user PATH, removed %d duplicates", removed)
	return nil
}

// normalizePath removes trailing slashes and trims whitespace.
func normalizePath(p string) string {
	p = strings.TrimSpace(p)
	p = strings.TrimSuffix(p, "\\")
	p = strings.TrimSuffix(p, "/")
	return strings.TrimSpace(p)
}
