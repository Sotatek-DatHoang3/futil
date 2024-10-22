package utils

import (
	"os"
	"testing"
)

func TestCheckFileExists(t *testing.T) {
	// create temporary file
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			return
		}
	}(tmpfile.Name())

	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{"Existing file", tmpfile.Name(), false},
		{"Non-existing file", "non_existing_file.txt", true},
		{"Stdin", "-", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckFileExists(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckFileExists() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOpenFileOrStdin(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{"Existing file", tmpfile.Name(), false},
		{"Non-existing file", "non_existing_file.txt", true},
		{"Stdin", "-", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader, err := OpenFileOrStdin(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenFileOrStdin() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				reader.Close()
			}
		})
	}
}
