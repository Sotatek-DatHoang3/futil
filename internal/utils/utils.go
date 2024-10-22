package utils

import (
	"fmt"
	"io"
	"os"
)

// HandleError print error and exit
func HandleError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}

// OpenFileOrStdin mở file hoặc trả về os.Stdin nếu filename là "-"
func OpenFileOrStdin(filename string) (io.ReadCloser, error) {
	if filename == "-" {
		return os.Stdin, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("không thể mở file: %v", err)
	}
	return file, nil
}

// CheckFileExists kiểm tra xem file có tồn tại không
func CheckFileExists(filename string) error {
	if filename == "-" {
		return nil // Cho phép đọc từ stdin
	}
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return fmt.Errorf("file không tồn tại: %s", filename)
	}
	if info.IsDir() {
		return fmt.Errorf("%s là một thư mục, không phải file", filename)
	}
	return nil
}
