package service

import (
	"fmt"
	"os"
	"time"
	"strconv"
	"task-cli/model"
	"task-cli/storage"
)

/// Handle the "add" command to add a new task
func HandleAdd() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide task description")
		return
	}

	description := os.Args[2]
	fmt.Println("Task to add: ", description)

	tasks, err := storage.ReadTasks()

	if err != nil {
		fmt.Println("Error reading tasks: ", err)
		return
	}

	newTask := model.Task{
		ID: GetNextTaskId(tasks),
		Description: description,
		Status:      model.StatusTodo,
		CreatedAt:   TimeStamp(),
		UpdatedAt:   TimeStamp(),
	}

	tasks = append(tasks, newTask)

	err = storage.WriteTasks(tasks)

	if err != nil {
		fmt.Println("Error writing tasks: ", err)
		return
	}

	fmt.Printf("Task added successfully (ID: %d)", newTask.ID)

}

/// Handle the update command to update a task's description by ID
func UpdateTask() {

	if len(os.Args) < 4 {
		fmt.Println("Please provide task ID and new description")
		return
	}

	taskID := os.Args[2]
	newDescription := os.Args[3]

	if taskID == "" {
		fmt.Println("Task ID cannot be empty")
		return
	}

	data, err := storage.ReadTasks()

	if err != nil {
		fmt.Println("Error reading tasks: ", err)
		return
	}

	tId, err := strconv.Atoi(taskID)

	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	found := false

	for idx, t := range data {
		if t.ID == int(tId) {
			data[idx].Description = newDescription
			data[idx].UpdatedAt = TimeStamp()
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Task with ID ", taskID, " not exists")
		return
	}

	err = storage.WriteTasks(data)

	if err != nil {
		fmt.Println("Error writing tasks: ", err)
		return
	}

	fmt.Println("Task updated successfully")
}

/// Handle the "mark-in-progress" command to update a task's status to "in progress" by ID
func MarkTaskInProgress() {
	
	if len(os.Args) < 3 {
		fmt.Println("Please provide task ID to mark in progress")
		return
	}

	taskID := os.Args[2]

	if taskID == "" {
		fmt.Println("Task ID cannot be empty")
		return
	}

	tId, err := strconv.Atoi(taskID)

	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	err = statusValidation(tId, model.StatusInProgress)

	if err != nil {
		fmt.Println("Error validating task status: ", err)
		return
	}

	err = UpdateTaskStatus(tId, model.StatusInProgress)

	if err != nil {
		fmt.Println("Error updating task status: ", err)
		return
	}

	fmt.Println("Task marked as in progress successfully")
}

/// Validate the current status of the task before updating to a new status
func statusValidation(id int, newStatus string) error {

	data, err := storage.ReadTasks()

	if err != nil {
		return fmt.Errorf("unable to read tasks: %v", err)
	}

	for _, t := range data {
		if t.ID == id {
			if t.Status == newStatus {
				return fmt.Errorf("Task with ID %d is already in status %s", id, newStatus)
			}
		}
	}

	return  nil
}

/// Validate the current status of the task before updating to a new status
func UpdateTaskStatus(taskId int, newStatus string) error {
	
	data, err := storage.ReadTasks()
	if err != nil {
		return err
	}

	found := false

	for idx, t := range data {
		if t.ID == taskId {
			data[idx].Status = newStatus
			data[idx].UpdatedAt = TimeStamp()
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("Task with ID %d not exists", taskId)
	}

	err = storage.WriteTasks(data)
	if err != nil {
		return err
	}

	fmt.Println("Task status updated successfully")
	return nil
}

/// Handle the delete command to delete a task by ID
func HandleDelete() {

	if len(os.Args) < 3 {
		fmt.Println("Please provide task ID to delete")
		return
	}

	taskID := os.Args[2]

	if taskID == "" {
		fmt.Println("Task ID cannot be empty")
		return
	}

	tId, err := strconv.Atoi(taskID)

	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	data, err := storage.ReadTasks()

	if err != nil {
		fmt.Println("Error reading tasks: ", err)
		return
	}

	foundIndex := -1

	for idx, t := range data {
		if t.ID == int(tId) {
			foundIndex = idx
			break
		}
	}

	if foundIndex == -1 {
		fmt.Println("Task with ID ", taskID, " not exists")
		return
	}

	data = append(data[:foundIndex], data[foundIndex+1:]...)
	err = storage.WriteTasks(data)
	if err != nil {
		fmt.Println("Error writing tasks: ", err)
		return
	}

	fmt.Println("Task deleted successfully")
}

/// Handle the "mark-done" command to update a task's status to "done" by ID
func MarkTaskDone() {
	
	if len(os.Args) < 3 {
		fmt.Println("Please provide task ID to mark done")
		return
	}

	taskID := os.Args[2]

	if taskID == "" {
		fmt.Println("Task ID cannot be empty")
		return
	}

	tId, err := strconv.Atoi(taskID)

	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	err = statusValidation(tId, model.StatusDone)

	if err != nil {
		fmt.Println("Error validating task status: ", err)
		return
	}

	err = UpdateTaskStatus(tId, model.StatusDone)

	if err != nil {
		fmt.Println("Error updating task status: ", err)
		return
	}

	fmt.Println("Task marked as done successfully")
}

/// This return time in RFC3339 format which is a standard format for date and time representation
func TimeStamp() string {
	return time.Now().Format(time.RFC3339)
}

/// GetNextTaskId returns the next available task ID based on the existing tasks
func GetNextTaskId(tasks []model.Task) int {
	if len(tasks) == 0 {
		return 1
	}

	return tasks[len(tasks)-1].ID + 1
}

/// Handle the help command to display usage instructions
func HandleHelp() {
	fmt.Println("Task Tracker CLI Application")
	fmt.Println("Usage:")
	fmt.Println("  add <description>          - Add a new task with the given description")
	fmt.Println("  list [status]              - List all tasks or filter by status (TODO, IN_PROGRESS, DONE)")
	fmt.Println("  update <id> <description>  - Update the description of a task by ID")
	fmt.Println("  mark-in-progress <id>      - Mark a task as in progress by ID")
	fmt.Println("  mark-done <id>             - Mark a task as done by ID")
	fmt.Println("  delete <id>                - Delete a task by ID")
	fmt.Println("  help                       - Show this help message")
}