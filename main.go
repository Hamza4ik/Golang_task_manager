package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID   int
	Text string
	Done bool
}

var tasks []Task

func addTask(tasks *[]Task, t Task) {
	*tasks = append(*tasks, t)
}

func listTasks(tasks []Task) {
	for _, t := range tasks {
		fmt.Printf("[%d] %s - done: %t\n", t.ID, t.Text, t.Done)
	}
}

func ReadInp() {
	fmt.Println("Writer Command: add, list, exit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		command := os.Args[0]

		switch command {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Длина parts меньше 2")
				continue
			}
			text := parts[1]
			addTask(&tasks, Task{ID: len(tasks) + 1, Text: text, Done: false})
			if err := saveTasks(tasks); err != nil {
				fmt.Println("Не удалось сохранить файл")
			}
		case "list":
			listTasks(tasks)
		}
	}
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

func loadTasks() ([]Task, error) {
	data, err := os.ReadFile("tasks.json")

	if os.IsNotExist(err) {
		return []Task{}, nil
	}
	if err != nil {
		return nil, err
	}

	var loaded []Task

	err = json.Unmarshal(data, &loaded)
	if err != nil {
		return nil, err
	}

	return loaded, nil
}

func main() {

	loaded, err := loadTasks()
	if err != nil {
		fmt.Println("Не удалось загрузить задачи, начинаю с пустого списка:", err)
		tasks = []Task{}
	} else {
		tasks = loaded
	}
	ReadInp()
}
