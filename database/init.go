package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DBPath = os.Getenv("HOME") + "/.pssh"
var DBFile = "hosts.db"

func InitDB() *sql.DB {
	_, err := os.Stat(DBPath)
	// check the DBPath exists
	if os.IsNotExist(err) {
		fmt.Printf("%s not exists, creating...\n", DBPath)
		err := os.MkdirAll(DBPath, 0755)
		if err != nil {
			fmt.Println("%v create failed: %v", DBPath, err)
			os.Exit(1)
		}
		fmt.Printf("%v create successfully.\n", DBPath)

	}
	// check the DBFile exists
	_, err = os.Stat(DBPath + "/" + DBFile)
	if os.IsNotExist(err) {
		fmt.Printf("%s not exists, creating...\n", DBFile)
		_, err := os.Create(DBPath + "/" + DBFile)
		if err != nil {
			fmt.Printf("%v create failed: %v\n", DBFile, err)
		}
		fmt.Printf("%v create successfully.\n", DBFile)
		fmt.Println("begin init db...")
		CreateTables()
	}

	db, err := sql.Open("sqlite3", DBPath)
	if err != nil {
		panic(err)
	}
	return db
}

func CreateTables() {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS hosts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		host VARCHER(50) NOT NULL UNIQUE,
		ip VARCHER(20) NOT NULL UNIQUE,
		port INT NOT NULL,
		username VARCHER(20) NOT NULL,
		password VARCHER(100) NOT NULL
	);`

	db := ConnDB()
	fmt.Println("init db table...")
	_, err := db.Exec(createTableSQL)
	if err != nil {
		fmt.Printf("failed to create table: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("init db table success.")

	defer db.Close()
}

func ConnDB() *sql.DB {
	db, err := sql.Open("sqlite3", filepath.Join(DBPath, DBFile))
	if err != nil {
		panic(err)
	}
	return db
}
