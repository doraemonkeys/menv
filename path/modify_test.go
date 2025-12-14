package path

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNormalizePath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "trailing backslash",
			path: "C:\\bin\\",
			want: "C:\\bin",
		},
		{
			name: "trailing forward slash",
			path: "C:/bin/",
			want: "C:/bin",
		},
		{
			name: "no trailing slash",
			path: "C:\\bin",
			want: "C:\\bin",
		},
		{
			name: "with leading spaces",
			path: "  C:\\bin",
			want: "C:\\bin",
		},
		{
			name: "with trailing spaces",
			path: "C:\\bin  ",
			want: "C:\\bin",
		},
		{
			name: "with both spaces and slash",
			path: "  C:\\bin\\  ",
			want: "C:\\bin",
		},
		{
			name: "multiple trailing slashes",
			path: "C:\\bin\\\\",
			want: "C:\\bin\\",
		},
		{
			name: "empty string",
			path: "",
			want: "",
		},
		{
			name: "only spaces",
			path: "   ",
			want: "",
		},
		{
			name: "path with spaces in name",
			path: "C:\\Program Files\\",
			want: "C:\\Program Files",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizePath(tt.path)
			if got != tt.want {
				t.Errorf("normalizePath() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestPathExists(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "existing directory",
			path: tempDir,
			want: true,
		},
		{
			name: "non-existing directory",
			path: filepath.Join(tempDir, "nonexistent"),
			want: false,
		},
		{
			name: "existing file",
			path: createTempFile(t, tempDir),
			want: true,
		},
		{
			name: "empty path",
			path: "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pathExists(tt.path)
			if got != tt.want {
				t.Errorf("pathExists(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}

func TestPathExistsWithEnvVar(t *testing.T) {
	tempDir := t.TempDir()
	os.Setenv("TEST_MENV_DIR", tempDir)
	defer os.Unsetenv("TEST_MENV_DIR")

	if !pathExists("$TEST_MENV_DIR") {
		t.Error("pathExists should expand environment variables with $VAR syntax")
	}
	if !pathExists("${TEST_MENV_DIR}") {
		t.Error("pathExists should expand environment variables with ${VAR} syntax")
	}
	if !pathExists("%TEST_MENV_DIR%") {
		t.Error("pathExists should expand environment variables with %VAR% syntax")
	}
}

func TestExpandWindowsEnv(t *testing.T) {
	os.Setenv("TEST_VAR1", "value1")
	os.Setenv("TEST_VAR2", "value2")
	defer os.Unsetenv("TEST_VAR1")
	defer os.Unsetenv("TEST_VAR2")

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "single variable",
			input: "%TEST_VAR1%",
			want:  "value1",
		},
		{
			name:  "multiple variables",
			input: "%TEST_VAR1%\\%TEST_VAR2%",
			want:  "value1\\value2",
		},
		{
			name:  "no variable",
			input: "C:\\bin",
			want:  "C:\\bin",
		},
		{
			name:  "undefined variable",
			input: "%UNDEFINED_VAR%",
			want:  "",
		},
		{
			name:  "mixed content",
			input: "C:\\%TEST_VAR1%\\bin",
			want:  "C:\\value1\\bin",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := expandWindowsEnv(tt.input)
			if got != tt.want {
				t.Errorf("expandWindowsEnv(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func createTempFile(t *testing.T, dir string) string {
	t.Helper()
	f, err := os.CreateTemp(dir, "test")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	return f.Name()
}
