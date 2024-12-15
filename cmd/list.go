package cmd

import (
	"fmt"
	"os"
	"personal/ssh-with-password/database"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the hosts in the database",
	Long:  `List all the hosts in the database`,
	Run: func(cmd *cobra.Command, args []string) {
		if all {
			HostFullList()
			return
		}
		HostList()
	},
}

var all bool

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "show hosts with all info")
}

func HostList() {
	db := database.ConnDB()
	defer db.Close()

	rows, err := db.Query("SELECT host, ip FROM hosts")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Host\tIp")
	for rows.Next() {
		var host, ip string
		err = rows.Scan(&host, &ip)
		if err != nil {
			fmt.Println("HostList query host failed, err:", err)
			return
		}
		fmt.Printf("%s\t%s\n", host, ip)
	}
	fmt.Println()
}

func HostFullList() {
	db := database.ConnDB()
	defer db.Close()

	rows, err := db.Query("SELECT host,ip,port,username FROM hosts")
	if err != nil {
		fmt.Println(err)
		return
	}

	// create a tabwriter
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	// print the header
	fmt.Fprintln(writer, "Host\t\tIP\t\tPort\t\tUsername")

	for rows.Next() {
		var host, ip, port, username string
		err = rows.Scan(&host, &ip, &port, &username)
		if err != nil {
			fmt.Println("HostFullList query host failed, err:", err)
			return
		}
		fmt.Fprintf(writer, "%s\t\t%s\t\t%s\t\t%s\n", host, ip, port, username)
	}
	// flush the writer
	writer.Flush()
	fmt.Println()
}
