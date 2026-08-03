package storage

import (
	"encoding/json"
	"os"

	"github.com/Hamza4ik/Golang_task_manager/internal/task"
)

func SaveTasks(tasks []task.Task, path string) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func LoadTasks(path string) ([]task.Task, error) {
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return []task.Task{}, nil
	}
	if err != nil {
		return nil, err
	}

	var loaded []task.Task
	err = json.Unmarshal(data, &loaded)
	if err != nil {
		return nil, err
	}
	return loaded, nil
}
