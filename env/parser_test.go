package env

import (
	"testing"

	"github.com/doraemonkeys/doraemon"
)

func TestParseEnvFile(t *testing.T) {
	tests := []struct {
		name      string
		content   []byte
		startWith string
		want      []doraemon.Pair[string, string]
		wantErr   bool
	}{
		{
			name:      "simple key-value",
			content:   []byte("KEY=value"),
			startWith: "",
			want:      []doraemon.Pair[string, string]{{First: "KEY", Second: "value"}},
			wantErr:   false,
		},
		{
			name:      "multiple lines",
			content:   []byte("KEY1=value1\nKEY2=value2"),
			startWith: "",
			want: []doraemon.Pair[string, string]{
				{First: "KEY1", Second: "value1"},
				{First: "KEY2", Second: "value2"},
			},
			wantErr: false,
		},
		{
			name:      "with comments",
			content:   []byte("# this is a comment\nKEY=value"),
			startWith: "",
			want:      []doraemon.Pair[string, string]{{First: "KEY", Second: "value"}},
			wantErr:   false,
		},
		{
			name:      "empty lines",
			content:   []byte("\n\nKEY=value\n\n"),
			startWith: "",
			want:      []doraemon.Pair[string, string]{{First: "KEY", Second: "value"}},
			wantErr:   false,
		},
		{
			name:      "with export prefix",
			content:   []byte("export KEY=value\nother=skip"),
			startWith: "export",
			want:      []doraemon.Pair[string, string]{{First: "KEY", Second: "value"}},
			wantErr:   false,
		},
		{
			name:      "trim spaces",
			content:   []byte("  KEY  =  value  "),
			startWith: "",
			want:      []doraemon.Pair[string, string]{{First: "KEY", Second: "value"}},
			wantErr:   false,
		},
		{
			name:      "missing equals",
			content:   []byte("KEYVALUE"),
			startWith: "",
			want:      nil,
			wantErr:   true,
		},
		{
			name:      "multiple equals error",
			content:   []byte("KEY=val=ue"),
			startWith: "",
			want:      nil,
			wantErr:   true,
		},
		{
			name:      "empty content",
			content:   []byte(""),
			startWith: "",
			want:      []doraemon.Pair[string, string]{},
			wantErr:   false,
		},
		{
			name:      "only comments",
			content:   []byte("# comment1\n# comment2"),
			startWith: "",
			want:      []doraemon.Pair[string, string]{},
			wantErr:   false,
		},
		{
			name:      "windows line endings",
			content:   []byte("KEY1=value1\r\nKEY2=value2"),
			startWith: "",
			want: []doraemon.Pair[string, string]{
				{First: "KEY1", Second: "value1"},
				{First: "KEY2", Second: "value2"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseEnvFile(tt.content, tt.startWith)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseEnvFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got) != len(tt.want) {
					t.Errorf("ParseEnvFile() got %d pairs, want %d", len(got), len(tt.want))
					return
				}
				for i := range got {
					if got[i].First != tt.want[i].First || got[i].Second != tt.want[i].Second {
						t.Errorf("ParseEnvFile()[%d] = {%s, %s}, want {%s, %s}",
							i, got[i].First, got[i].Second, tt.want[i].First, tt.want[i].Second)
					}
				}
			}
		})
	}
}
