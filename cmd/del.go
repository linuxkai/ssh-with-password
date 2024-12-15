package cmd

import (
	"fmt"
	"personal/ssh-with-password/database"
	"personal/ssh-with-password/utils"

	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a host from the database",
	Long: `Delete a host from the database.
	
	Example:
	del ip 192.168.1.1
	del host myhost`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("Error: Only one host can be deleted at a time")
			return
		}
		if len(args) == 0 {
			fmt.Println("Error: Please provide a host to delete")
			return
		}

		DelHost(args[0])
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}

func DelHost(delStr string) {
	var delFlag = "host"
	isIP, _ := utils.ValidateIP(delStr)
	if isIP {
		delFlag = "ip"
	}
	db := database.ConnDB()
	defer db.Close()

	// delete host
	sql := fmt.Sprintf("delete from hosts where %s = ?", delFlag)
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println("DelHost prepare host err: ", err)
		return
	}

	res, err := stmt.Exec(delStr)
	if err != nil {
		fmt.Println("DelHost exec sql ", err)
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	if affect > 0 {
		fmt.Println("delete host successfully.")
	}
}
