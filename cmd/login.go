package cmd

import (
	"bytes"
	"database/sql"
	"fmt"
	"os"
	"personal/ssh-with-password/database"
	"personal/ssh-with-password/utils"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

type Host struct {
	Ip       string
	Port     string
	Username string
	Password string
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to host",
	Long:  `login to host`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			host_info, err := getHost(args[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			password, err := utils.Decrypt(host_info.Password)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			SshHost(host_info.Ip, host_info.Port, host_info.Username, password)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func getHost(host string) (Host, error) {
	host_info := Host{}
	if host == "" {
		return host_info, fmt.Errorf("host name can not be empty")
	}

	var query_flag = "host"
	isIP, _ := utils.ValidateIP(host)
	if isIP {
		query_flag = "ip"
	}
	db := database.ConnDB()
	defer db.Close()

	query_sql := fmt.Sprintf("select ip, port, username, password from hosts where %s = ?", query_flag)
	err := db.QueryRow(query_sql, host).Scan(&host_info.Ip, &host_info.Port, &host_info.Username, &host_info.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found.")
			fmt.Println(err)
		} else {
			fmt.Println("Failed to query row: %s", err)
		}
		return host_info, fmt.Errorf("host name not found")
	}

	return host_info, nil
}

func SshHost(ip, port, username, password string) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 忽略主机密钥验证（生产环境请替换为更安全的实现）
	}

	// connect the host
	addr := fmt.Sprintf("%s:%s", ip, port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		fmt.Println("Failed to dial: %s", err)
	}
	defer client.Close()

	// create a new session
	session, err := client.NewSession()
	if err != nil {
		fmt.Println("Failed to create session: %s", err)
	}
	defer session.Close()

	// excute a command
	var stdout, stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr

	//create a terminal
	// set terminal mode
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // 关闭回显
		ssh.TTY_OP_ISPEED: 14400, // 设置传输速率
		ssh.TTY_OP_OSPEED: 14400,
	}

	// request a pseudo terminal
	err = session.RequestPty("linux", 32, 160, modes)
	if err != nil {
		fmt.Println(err)
	}
	// set stdout and stdin
	session.Stdout = os.Stdout
	session.Stdin = os.Stdin
	session.Stderr = os.Stderr

	session.Shell() // start shell
	session.Wait()  // wait for the command to finish
}
