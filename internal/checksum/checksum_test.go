package checksum

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"io"
	"strings"
	"testing"
)

func TestCalculateAndPrintChecksum(t *testing.T) {
	input := "Hello, World!"
	tests := []struct {
		name     string
		hashFunc func() hash.Hash
		want     string
	}{
		{"MD5", md5.New, "65a8e27d8879283831b664bd8b7f0ad4"},
		{"SHA1", sha1.New, "0a0a9f2a6772942557ab5355d76af442f8f65e01"},
		{"SHA256", sha256.New, "dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(input)
			h := tt.hashFunc()

			_, err := io.Copy(h, reader)
			if err != nil {
				t.Fatalf("Không thể tính toán hash: %v", err)
			}

			got := hex.EncodeToString(h.Sum(nil))
			if got != tt.want {
				t.Errorf("calculateAndPrintChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}
