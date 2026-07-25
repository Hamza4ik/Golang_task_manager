package storage

import (
	"encoding/json"
	"os"

	"github.com/Hamza4ik/Golang_task_manager/internal/task"
)

func SaveTasks(tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

func LoadTasks() ([]task.Task, error) {
	data, err := os.ReadFile("tasks.json")
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
