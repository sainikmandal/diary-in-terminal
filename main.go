package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n📔 Welcome to Your Diary")
		fmt.Println("1. Insert new entry")
		fmt.Println("2. View entries")
		fmt.Println("3. Delete entry")
		fmt.Println("4. Export entries")
		fmt.Println("5. Exit")
		fmt.Print("Select an option (1-5): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("👉 Insert new entry (not implemented yet)")
		case "2":
			fmt.Println("📖 View entries (not implemented yet)")
		case "3":
			fmt.Println("❌ Delete entry (not implemented yet)")
		case "4":
			fmt.Println("📤 Export entries (not implemented yet)")
		case "5":
			fmt.Println("👋 Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("❗ Invalid choice. Try again.")
		}
	}
}
