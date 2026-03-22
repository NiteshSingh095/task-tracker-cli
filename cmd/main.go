package main

import (
	"fmt"
	"os"
	"task-cli/model"
	"task-cli/storage"
	"task-cli/service"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide command")
		return
	}

	err := storage.EnsureTaskFileExists()
	if err != nil {
		fmt.Println("Error ensuring task file exists: ", err)
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		service.HandleAdd()

	case "list":
		handleList()

	case "delete":
		service.HandleDelete()

	case "update":
		service.UpdateTask()

	case "mark-in-progress":
		service.MarkTaskInProgress()

	case "mark-done":
		service.MarkTaskDone()

	case "help":
		service.HandleHelp()

	default:
		fmt.Println("Unknown Command")
	}
}

/// Handle the List command to display all tasks or filter by status if provided
func handleList() {
	fmt.Println("Handling List Command")

	task, err := storage.ReadTasks()

	if err != nil {
		fmt.Println("Error reading tasks: ", err)
		return
	}

	if len(os.Args) == 2 {
		printTaskDetails(task)
		return
	}

	status := os.Args[2]

	if status == "" {
		fmt.Println("Status cannot be empty")
		return
	}

	if status != model.StatusTodo && status != model.StatusInProgress && status != model.StatusDone {
		fmt.Println("Invalid status. Allowed values are TODO, IN_PROGRESS, DONE")
		return
	}

	var filteredTask []model.Task

	for _, t := range task {
		if t.Status == status {
			filteredTask = append(filteredTask, t)
		}
	}

	printTaskDetails(filteredTask)

}

/// Utility function to print task details in a readable format
func printTaskDetails(task []model.Task) {

	if len(task) == 0 {
		fmt.Println("No tasks found")
		return
	}

	for _, t := range task {
		fmt.Printf("[%d] %s (%s)\n", t.ID, t.Description, t.Status)
	}
}

