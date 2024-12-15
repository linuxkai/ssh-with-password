package cmd

import (
	"personal/ssh-with-password/database"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init database.",
	Long:  "init database and create tables",
	Run: func(cmd *cobra.Command, args []string) {
		database.InitDB()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
