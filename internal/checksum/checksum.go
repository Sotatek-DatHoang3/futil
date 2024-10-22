package checksum

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"futil/internal/utils"
	"github.com/spf13/cobra"
	"hash"
	"io"
)

func NewChecksumCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "checksum",
		Short: "Print checksum of file",
		Long:  `Print the checksum of the specified file using various algorithms.`,
		RunE:  runChecksum,
	}

	cmd.Flags().StringP("file", "f", "", "the input file")
	_ = cmd.MarkFlagRequired("file")
	cmd.Flags().Bool("md5", false, "use MD5 algorithm")
	cmd.Flags().Bool("sha1", false, "use SHA1 algorithm")
	cmd.Flags().Bool("sha256", false, "use SHA256 algorithm")

	return cmd
}

func runChecksum(cmd *cobra.Command, args []string) error {
	filename, _ := cmd.Flags().GetString("file")
	useMD5, _ := cmd.Flags().GetBool("md5")
	useSHA1, _ := cmd.Flags().GetBool("sha1")
	useSHA256, _ := cmd.Flags().GetBool("sha256")

	if err := utils.CheckFileExists(filename); err != nil {
		return err
	}

	reader, err := utils.OpenFileOrStdin(filename)
	if err != nil {
		return err
	}
	defer func(reader io.ReadCloser) {
		errC := reader.Close()
		if errC != nil {
			return
		}
	}(reader)

	if useMD5 {
		if err := calculateAndPrintChecksum(reader, md5.New()); err != nil {
			return err
		}
	}
	if useSHA1 {
		if err := calculateAndPrintChecksum(reader, sha1.New()); err != nil {
			return err
		}
	}
	if useSHA256 {
		if err := calculateAndPrintChecksum(reader, sha256.New()); err != nil {
			return err
		}
	}

	if !useMD5 && !useSHA1 && !useSHA256 {
		return fmt.Errorf("error: Please specify at least one checksum algorithm")
	}

	return nil
}

func calculateAndPrintChecksum(r io.Reader, h hash.Hash) error {
	if _, err := io.Copy(h, r); err != nil {
		return fmt.Errorf("error calculating checksum: %v", err)
	}
	fmt.Printf("%x\n", h.Sum(nil))
	return nil
}
