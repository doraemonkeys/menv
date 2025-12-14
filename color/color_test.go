package color

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestSprintf(t *testing.T) {
	tests := []struct {
		name   string
		color  string
		format string
		args   []any
		want   string
	}{
		{
			name:   "red text",
			color:  Red,
			format: "error",
			args:   nil,
			want:   Red + "error" + Reset,
		},
		{
			name:   "green with format",
			color:  Green,
			format: "success: %s",
			args:   []any{"done"},
			want:   Green + "success: done" + Reset,
		},
		{
			name:   "cyan with number",
			color:  Cyan,
			format: "count: %d",
			args:   []any{42},
			want:   Cyan + "count: 42" + Reset,
		},
		{
			name:   "bold yellow",
			color:  BoldYellow,
			format: "warning",
			args:   nil,
			want:   BoldYellow + "warning" + Reset,
		},
		{
			name:   "empty string",
			color:  Blue,
			format: "",
			args:   nil,
			want:   Blue + "" + Reset,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sprintf(tt.color, tt.format, tt.args...)
			if got != tt.want {
				t.Errorf("Sprintf() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestPrintColoredLn(t *testing.T) {
	tests := []struct {
		name   string
		color  string
		format string
		args   []any
	}{
		{
			name:   "simple message",
			color:  Green,
			format: "hello",
			args:   nil,
		},
		{
			name:   "formatted message",
			color:  Red,
			format: "error: %s",
			args:   []any{"failed"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			printColoredLn(tt.color, tt.format, tt.args...)

			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			buf.ReadFrom(r)
			output := buf.String()

			if !strings.HasPrefix(output, tt.color) {
				t.Errorf("output should start with color code")
			}
			if !strings.HasSuffix(output, Reset+"\n") {
				t.Errorf("output should end with reset code and newline")
			}
		})
	}
}

func TestSuccess(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Success("test %s", "message")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "test message") {
		t.Errorf("Success() output should contain message, got %q", output)
	}
	if !strings.HasPrefix(output, Green) {
		t.Errorf("Success() should use green color")
	}
}

func TestError(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Error("error %d", 404)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "error 404") {
		t.Errorf("Error() output should contain message, got %q", output)
	}
	if !strings.HasPrefix(output, Red) {
		t.Errorf("Error() should use red color")
	}
}

func TestWarning(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Warning("warn")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "warn") {
		t.Errorf("Warning() output should contain message, got %q", output)
	}
	if !strings.HasPrefix(output, Yellow) {
		t.Errorf("Warning() should use yellow color")
	}
}

func TestInfo(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Info("info")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "info") {
		t.Errorf("Info() output should contain message, got %q", output)
	}
	if !strings.HasPrefix(output, Cyan) {
		t.Errorf("Info() should use cyan color")
	}
}

func TestHighlight(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Highlight("highlight")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "highlight") {
		t.Errorf("Highlight() output should contain message, got %q", output)
	}
	if !strings.HasPrefix(output, Magenta) {
		t.Errorf("Highlight() should use magenta color")
	}
}
