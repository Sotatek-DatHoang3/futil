package main

import (
	"fmt"
	"futil/internal/utils"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "futil",
	Short: "File Utility",
	Long:  `File Utility is a CLI tool for file operations.`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("futil v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		utils.HandleError(err)
	}
}
