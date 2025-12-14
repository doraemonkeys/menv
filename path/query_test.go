package path

import (
	"reflect"
	"testing"
)

func TestSplitAndCleanPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want []string
	}{
		{
			name: "simple paths",
			path: "C:\\bin;D:\\tools",
			want: []string{"C:\\bin", "D:\\tools"},
		},
		{
			name: "with empty entries",
			path: "C:\\bin;;D:\\tools",
			want: []string{"C:\\bin", "D:\\tools"},
		},
		{
			name: "with spaces",
			path: " C:\\bin ; D:\\tools ",
			want: []string{"C:\\bin", "D:\\tools"},
		},
		{
			name: "trailing semicolon",
			path: "C:\\bin;D:\\tools;",
			want: []string{"C:\\bin", "D:\\tools"},
		},
		{
			name: "single path",
			path: "C:\\bin",
			want: []string{"C:\\bin"},
		},
		{
			name: "empty string",
			path: "",
			want: []string{},
		},
		{
			name: "only semicolons",
			path: ";;;",
			want: []string{},
		},
		{
			name: "path with spaces in name",
			path: "C:\\Program Files;D:\\My Apps",
			want: []string{"C:\\Program Files", "D:\\My Apps"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitAndCleanPath(tt.path)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitAndCleanPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
