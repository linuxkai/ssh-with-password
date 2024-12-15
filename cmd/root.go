package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"personal/ssh-with-password/database"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pssh",
	Short: "pssh is a command line tool for useing SSH to login host with password.",
	Long: `pssh is a command line tool for via SSH to login host with password.
    you can use it to login host with password and management multiple hosts.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(args[0])
			if args[0] == "init" {
				return
			} else {
				_, err := os.Stat(filepath.Join(database.DBPath, database.DBFile))
				if os.IsNotExist(err) {
					fmt.Println("WARNING: First run , please use 'pssh init' to init db")
					os.Exit(1)
				}
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
