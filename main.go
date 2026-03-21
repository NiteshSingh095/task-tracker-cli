package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("All OS Arguments : ", os.Args)

	if(len(os.Args) < 2) {
		fmt.Println("Please provide command")
		return
	}

	cmd := os.Args[1]
	
	switch cmd {
	case "add":
		handleAdd()

	case "list":
		handleList()

	case "delete":
		handleDelete()

	default:
		fmt.Println("Unknown Command")
	}
}

func handleAdd() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide task description")
		return
	}

	description := os.Args[2]
	fmt.Println("Task to add: ", description)
}

func handleList() {
	fmt.Println("Handling List Command")
}

func handleDelete() {
	
	if len(os.Args) < 3 {
		fmt.Println("Please provide task ID to delete")
		return
	}

	taskID := os.Args[2]
	fmt.Println("Task ID to delete: ", taskID)
}