package path

import "testing"

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
