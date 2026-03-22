package storage

import (
	"os"
	"encoding/json"
	"task-cli/model"
)

const taskFile = "tasks.json"

/// Ensure the task file exists, if not create it with an empty array
func EnsureTaskFileExists() error {

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

/// Read tasks from the JSON file and return as a slice of Task
func ReadTasks() ([]model.Task, error) {
	data, err := os.ReadFile(taskFile)

	if err != nil {
		return nil, err
	}

	var tasks []model.Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

/// Write the given slice of Task to the JSON file
func WriteTasks(tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(taskFile, data, 0644)
}
