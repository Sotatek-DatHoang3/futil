package linecount

import (
	"bufio"
	"fmt"
	"io"

	"futil/internal/utils"
	"github.com/spf13/cobra"
)

func NewLinecountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "linecount",
		Short: "Print line count of file",
		Long:  `Print the number of lines in the specified file or from stdin.`,
		RunE:  runLinecount,
	}

	cmd.Flags().StringP("file", "f", "", "the input file")
	_ = cmd.MarkFlagRequired("file")
	return cmd
}

func runLinecount(cmd *cobra.Command, args []string) error {
	filename, _ := cmd.Flags().GetString("file")

	if err := utils.CheckFileExists(filename); err != nil {
		return err
	}

	reader, err := utils.OpenFileOrStdin(filename)
	if err != nil {
		return err
	}
	defer reader.Close()

	count, err := countLines(reader)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	fmt.Printf("line count: %d \n", count)
	return nil
}

func countLines(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	return lineCount, scanner.Err()
}
