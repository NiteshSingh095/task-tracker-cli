package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

const taskFile = "tasks.json"

func main() {
	fmt.Println("All OS Arguments : ", os.Args)

	if len(os.Args) < 2 {
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

	case "update":
		updateTask()

	default:
		fmt.Println("Unknown Command")
	}
}

func timeStamp() string {
	return time.Now().Format(time.RFC3339)
}

func updateTask() {

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

	data, err := readTasks()

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
			data[idx].UpdatedAt = timeStamp()
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Task with ID ", taskID, " not exists")
		return
	}

	err = WriteTasks(data)

	if err != nil {
		fmt.Println("Error writing tasks: ", err)
		return
	}

	fmt.Println("Task updated successfully")
}

func handleAdd() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide task description")
		return
	}

	description := os.Args[2]
	fmt.Println("Task to add: ", description)

	tasks, err := readTasks()

	if err != nil {
		fmt.Println("Error reading tasks: ", err)
		return
	}

	newTask := Task{
		ID:          getNextTaskId(tasks),
		Description: description,
		Status:      "TODO",
		CreatedAt:   timeStamp(),
		UpdatedAt:   timeStamp(),
	}

	tasks = append(tasks, newTask)

	err = WriteTasks(tasks)

	if err != nil {
		fmt.Println("Error writing tasks: ", err)
		return
	}

	fmt.Println("Task added successfully : ", newTask.ID)

}

func handleList() {
	fmt.Println("Handling List Command")

	task, err := readTasks()

	if err != nil {
		fmt.Println("Error reading tasks: ", err)
		return
	}

	for _, t := range task {
		fmt.Printf("[%d] %s (%s)\n", t.ID, t.Description, t.Status)
	}

}

func handleDelete() {

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

	data, err := readTasks()

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
	err = WriteTasks(data)
	if err != nil {
		fmt.Println("Error writing tasks: ", err)
		return
	}

	fmt.Println("Task deleted successfully")
}

func getNextTaskId(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}

	return tasks[len(tasks)-1].ID + 1
}

func readTasks() ([]Task, error) {
	data, err := os.ReadFile(taskFile)

	if err != nil {
		return nil, err
	}

	var tasks []Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func WriteTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

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
