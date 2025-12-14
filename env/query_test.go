package env

import (
	"runtime"
	"testing"
)

func TestParseRegLine(t *testing.T) {
	tests := []struct {
		name    string
		line    string
		wantKey string
		wantVal string
		wantNil bool
	}{
		{
			name:    "REG_SZ simple",
			line:    "    GOPATH    REG_SZ    C:\\Go",
			wantKey: "GOPATH",
			wantVal: "C:\\Go",
		},
		{
			name:    "REG_EXPAND_SZ",
			line:    "    Path    REG_EXPAND_SZ    %USERPROFILE%\\bin",
			wantKey: "Path",
			wantVal: "%USERPROFILE%\\bin",
		},
		{
			name:    "value with spaces",
			line:    "    JAVA_HOME    REG_SZ    C:\\Program Files\\Java\\jdk",
			wantKey: "JAVA_HOME",
			wantVal: "C:\\Program Files\\Java\\jdk",
		},
		{
			name:    "unsupported type REG_DWORD",
			line:    "    COUNT    REG_DWORD    0x1",
			wantNil: true,
		},
		{
			name:    "too few fields",
			line:    "    ONLY_KEY",
			wantNil: true,
		},
		{
			name:    "empty line",
			line:    "",
			wantNil: true,
		},
		{
			name:    "only spaces",
			line:    "     ",
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseRegLine(tt.line)
			if tt.wantNil {
				if got != nil {
					t.Errorf("parseRegLine() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Errorf("parseRegLine() = nil, want {%s, %s}", tt.wantKey, tt.wantVal)
				return
			}
			if got.Key != tt.wantKey {
				t.Errorf("parseRegLine().Key = %s, want %s", got.Key, tt.wantKey)
			}
			if got.Value != tt.wantVal {
				t.Errorf("parseRegLine().Value = %s, want %s", got.Value, tt.wantVal)
			}
		})
	}
}

func TestParseRegOutput(t *testing.T) {
	tests := []struct {
		name   string
		output string
		want   []EnvVar
	}{
		{
			name: "typical output",
			output: `HKEY_CURRENT_USER\Environment
    GOPATH    REG_SZ    C:\Go
    JAVA_HOME    REG_SZ    C:\Java

`,
			want: []EnvVar{
				{Key: "GOPATH", Value: "C:\\Go"},
				{Key: "JAVA_HOME", Value: "C:\\Java"},
			},
		},
		{
			name: "with REG_EXPAND_SZ",
			output: `HKEY_CURRENT_USER\Environment
    Path    REG_EXPAND_SZ    %USERPROFILE%\bin
    TEMP    REG_EXPAND_SZ    %USERPROFILE%\Temp
`,
			want: []EnvVar{
				{Key: "Path", Value: "%USERPROFILE%\\bin"},
				{Key: "TEMP", Value: "%USERPROFILE%\\Temp"},
			},
		},
		{
			name: "mixed types filters non-string",
			output: `HKEY_CURRENT_USER\Environment
    NAME    REG_SZ    value
    COUNT    REG_DWORD    0x1
`,
			want: []EnvVar{
				{Key: "NAME", Value: "value"},
			},
		},
		{
			name:   "empty output",
			output: "",
			want:   nil,
		},
		{
			name: "only registry path",
			output: `HKEY_CURRENT_USER\Environment
`,
			want: nil,
		},
		{
			name: "sorted alphabetically",
			output: `HKEY_CURRENT_USER\Environment
    ZEBRA    REG_SZ    z
    APPLE    REG_SZ    a
    MANGO    REG_SZ    m
`,
			want: []EnvVar{
				{Key: "APPLE", Value: "a"},
				{Key: "MANGO", Value: "m"},
				{Key: "ZEBRA", Value: "z"},
			},
		},
		{
			name: "case insensitive sort",
			output: `HKEY_CURRENT_USER\Environment
    Zebra    REG_SZ    z
    apple    REG_SZ    a
`,
			want: []EnvVar{
				{Key: "apple", Value: "a"},
				{Key: "Zebra", Value: "z"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseRegOutput(tt.output)
			if len(got) != len(tt.want) {
				t.Errorf("parseRegOutput() got %d items, want %d", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i].Key != tt.want[i].Key || got[i].Value != tt.want[i].Value {
					t.Errorf("parseRegOutput()[%d] = {%s, %s}, want {%s, %s}",
						i, got[i].Key, got[i].Value, tt.want[i].Key, tt.want[i].Value)
				}
			}
		})
	}
}

func TestListUser(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("skipping test on non-Windows OS")
	}
	envVars, err := ListUser()
	if err != nil {
		t.Errorf("ListUser() error = %v", err)
		return
	}
	if len(envVars) == 0 {
		t.Log("ListUser() returned empty list (may be expected on some systems)")
	}
}

func TestListSystem(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("skipping test on non-Windows OS")
	}
	envVars, err := ListSystem()
	if err != nil {
		t.Errorf("ListSystem() error = %v", err)
		return
	}
	if len(envVars) == 0 {
		t.Error("ListSystem() returned empty list, expected at least some system env vars")
	}
}

func TestGetUser(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("skipping test on non-Windows OS")
	}
	val, err := GetUser("MENV_TEST_NONEXISTENT_KEY_12345")
	if err != nil {
		t.Errorf("GetUser() error = %v", err)
		return
	}
	if val != "" {
		t.Errorf("GetUser() for non-existent key = %s, want empty", val)
	}
}

func TestGetSystem(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("skipping test on non-Windows OS")
	}
	val, err := GetSystem("MENV_TEST_NONEXISTENT_KEY_12345")
	if err != nil {
		t.Errorf("GetSystem() error = %v", err)
		return
	}
	if val != "" {
		t.Errorf("GetSystem() for non-existent key = %s, want empty", val)
	}
}
