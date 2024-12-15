package cmd

import (
	"fmt"
	"personal/ssh-with-password/database"
	"personal/ssh-with-password/utils"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add host to database.",
	Long:  `add host to database.`,
	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]
		ip := args[1]
		port := args[2]
		username := args[3]
		password := args[4]
		// validate ip and port format
		isIP, err := utils.ValidateIP(ip)
		if !isIP {
			fmt.Println(err)
			return
		}
		isPort, err := utils.ValidatePort(port)
		if !isPort {
			fmt.Println(err)
			return
		}
		AddHost(host, ip, port, username, password)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func AddHost(host, ip, port, username, password string) {
	db := database.ConnDB()
	defer db.Close()

	// check the host and ip whether exists
	query_sql := "SELECT host,ip,port,username,password FROM hosts WHERE host=? or ip=?"
	rows, err := db.Query(query_sql, host, ip)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		fmt.Println("host already exists.")
		return
	}

	sql := "INSERT INTO hosts(host, ip, port, username, password) VALUES(?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	// password encrypt
	encryptPassword, err := utils.Encrypt(password)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(host, ip, port, username, encryptPassword)
	if err != nil {
		panic(err)
	}
	fmt.Println("add host successfully.")
}
