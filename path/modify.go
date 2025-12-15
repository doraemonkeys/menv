package path

import (
	"errors"
	"os"
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

// Remove removes a path from the PATH environment variable.
func Remove(remove string, sys bool) error {
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

	remove = normalizePath(remove)
	removeNorm := strings.ToLower(remove)

	var newPath string
	found := false
	for _, p := range paths {
		pNorm := strings.ToLower(normalizePath(p))
		if pNorm == removeNorm {
			found = true
			continue
		}
		newPath += p + ";"
	}

	if !found {
		color.Warning("path not found: %s", remove)
		return nil
	}

	scope := "user"
	target := "User"
	if sys {
		scope = "system"
		target = "Machine"
	}

	cmd := exec.Command("powershell", "[System.Environment]::SetEnvironmentVariable(\"Path\", \""+newPath+"\", \""+target+"\")")
	if err := cmd.Run(); err != nil {
		return err
	}
	color.Success("removed %s [%s PATH]", remove, scope)
	return nil
}

// CleanResult contains the preview of paths to be cleaned.
type CleanResult struct {
	Duplicates []string
	Invalid    []string
	NewPath    string
}

// PreviewClean analyzes PATH and returns what would be cleaned.
func PreviewClean(sys bool) (CleanResult, error) {
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
		return CleanResult{}, err
	}

	seen := make(map[string]bool, len(paths))
	var result CleanResult

	for _, p := range paths {
		pNorm := strings.ToLower(normalizePath(p))

		if seen[pNorm] {
			result.Duplicates = append(result.Duplicates, p)
			continue
		}
		seen[pNorm] = true

		if !pathExists(p) {
			result.Invalid = append(result.Invalid, p)
			continue
		}

		result.NewPath += p + ";"
	}

	return result, nil
}

// ApplyClean applies the cleaned PATH to the registry.
func ApplyClean(newPath string, sys bool) error {
	scope := "user"
	target := "User"
	if sys {
		scope = "system"
		target = "Machine"
	}

	cmd := exec.Command("powershell", "[System.Environment]::SetEnvironmentVariable(\"Path\", \""+newPath+"\", \""+target+"\")")
	if err := cmd.Run(); err != nil {
		return err
	}
	color.Success("cleaned %s PATH", scope)
	return nil
}

func pathExists(p string) bool {
	p = expandPath(p)
	_, err := os.Stat(p)
	return err == nil
}

func expandPath(p string) string {
	p = os.ExpandEnv(p)
	p = os.Expand(p, os.Getenv)
	return expandWindowsEnv(p)
}

// expandWindowsEnv expands Windows-style %VAR% environment variables.
func expandWindowsEnv(p string) string {
	for {
		start := strings.Index(p, "%")
		if start == -1 {
			break
		}
		end := strings.Index(p[start+1:], "%")
		if end == -1 {
			break
		}
		end += start + 1
		varName := p[start+1 : end]
		value := os.Getenv(varName)
		p = p[:start] + value + p[end+1:]
	}
	return p
}

// normalizePath removes trailing slashes and trims whitespace.
func normalizePath(p string) string {
	p = strings.TrimSpace(p)
	p = strings.TrimSuffix(p, "\\")
	p = strings.TrimSuffix(p, "/")
	return strings.TrimSpace(p)
}

// InvalidPath represents an invalid path entry with its index.
type InvalidPath struct {
	Index int
	Path  string
}

// Check finds invalid paths in PATH that don't exist on filesystem.
func Check(sys bool) ([]InvalidPath, error) {
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
		return nil, err
	}

	var invalid []InvalidPath
	for i, p := range paths {
		if !pathExists(p) {
			invalid = append(invalid, InvalidPath{Index: i + 1, Path: p})
		}
	}
	return invalid, nil
}

// RemoveInvalidPaths removes specified invalid paths from PATH.
func RemoveInvalidPaths(paths []InvalidPath, sys bool) error {
	var (
		currentPaths []string
		err          error
	)

	if sys {
		currentPaths, err = QuerySystemPath()
	} else {
		currentPaths, err = QueryUserPath()
	}
	if err != nil {
		return err
	}

	toRemove := make(map[string]bool, len(paths))
	for _, p := range paths {
		toRemove[strings.ToLower(normalizePath(p.Path))] = true
	}

	var newPath string
	for _, p := range currentPaths {
		pNorm := strings.ToLower(normalizePath(p))
		if toRemove[pNorm] {
			continue
		}
		newPath += p + ";"
	}

	scope := "user"
	target := "User"
	if sys {
		scope = "system"
		target = "Machine"
	}

	cmd := exec.Command("powershell", "[System.Environment]::SetEnvironmentVariable(\"Path\", \""+newPath+"\", \""+target+"\")")
	if err := cmd.Run(); err != nil {
		return err
	}
	color.Success("removed %d invalid paths from %s PATH", len(paths), scope)
	return nil
}
