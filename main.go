package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sainikmandal/diary-in-terminal/auth"
	"github.com/sainikmandal/diary-in-terminal/db"
	_ "modernc.org/sqlite"
)

func main() {
	database, err := sql.Open("sqlite", "./diary.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	err = db.InitDB(database)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Terminal Diary")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Print("Select option: ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	var key []byte
	if choice == "1" {
		key, err = auth.Authenticate(database, password)
		if err != nil {
			log.Fatal("Authentication failed:", err)
		}
	} else {
		key, err = auth.Register(database, password)
		if err != nil {
			log.Fatal("Registration failed:", err)
		}
	}

	for {
		fmt.Println("\nCommands: insert | view | delete | exit")
		fmt.Print("Enter command: ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		switch cmd {
		case "insert":
			fmt.Print("Write your diary entry: ")
			entry, _ := reader.ReadString('\n')
			err := db.InsertEntry(database, key, entry)
			if err != nil {
				fmt.Println("Error inserting:", err)
			} else {
				fmt.Println("Entry saved.")
			}
		case "view":
			err := db.ViewEntries(database, key)
			if err != nil {
				fmt.Println("Error viewing entries:", err)
			}
		case "delete":
			fmt.Print("Enter entry ID to delete: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			err := db.DeleteEntry(database, idStr)
			if err != nil {
				fmt.Println("Delete failed:", err)
			} else {
				fmt.Println("Deleted.")
			}
		case "exit":
			return
		default:
			fmt.Println("Unknown command.")
		}
	}
}
