package linecount

import (
	"bytes"
	"testing"
)

func TestCountLines(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{"Empty string", "", 0, false},
		{"Single line", "Hello, World!", 1, false},
		{"Multiple lines", "Line 1\nLine 2\nLine 3", 3, false},
		{"Ends with newline", "Line 1\nLine 2\n", 2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bytes.NewBufferString(tt.input)
			got, err := countLines(reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("countLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("countLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
