package main

import (
	"fmt"
	"os"
	// "io/ioutil"
)

const taskFile = "tasks.json"

func main() {
	fmt.Println("All OS Arguments : ", os.Args)

	if(len(os.Args) < 2) {
		fmt.Println("Please provide command")
		return
	}

	err := ensureTaskFileExists()
	if err != nil {
		fmt.Println("Error ensuring task file exists: ", err)
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

	WriteTasks([]byte(`[{"test": "test task"}]`))
}

func handleList() {
	fmt.Println("Handling List Command")

	data, err := readTasks()

	if err != nil {
		fmt.Println("Error reading tasks: ", err)
		return
	}

	fmt.Println("Tasks: ", string(data))
}

func handleDelete() {
	
	if len(os.Args) < 3 {
		fmt.Println("Please provide task ID to delete")
		return
	}

	taskID := os.Args[2]
	fmt.Println("Task ID to delete: ", taskID)
}

func readTasks() ([]byte, error) {
	return os.ReadFile(taskFile)
}

func WriteTasks(data []byte) error {
	return os.WriteFile(taskFile, data, 0644)
}

func ensureTaskFileExists() error {

	_, err := os.Stat(taskFile)

	if os.IsNotExist(err) {

		file, err := os.Create(taskFile)

		if err != nil {
			return err
		}

		defer file.Close()

		_, err = file.Write([]byte("[]"))

		if err != nil {
			return err
		}
	}

	return nil
}