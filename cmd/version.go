package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of pssh",
	Long:  `Print the version number of pssh`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pssh version 0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
