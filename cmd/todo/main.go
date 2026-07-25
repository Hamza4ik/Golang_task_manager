package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Hamza4ik/Golang_task_manager/internal/storage"
	"github.com/Hamza4ik/Golang_task_manager/internal/task"
)

var tasks []task.Task

func ReadInp() {
	fmt.Println("Write command: add, list, exit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		command := parts[0]

		switch command {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Len parts less then 2")
				continue
			}
			text := parts[1]
			task.AddTask(&tasks, task.Task{ID: len(tasks) + 1, Text: text, Done: false})
			if err := storage.SaveTasks(tasks); err != nil {
				fmt.Println("Не удалось сохранить")
			}
		case "list":
			task.ListTasks(tasks)
		case "exit":
			return
		default:
			fmt.Println("Неузнаваемая команда")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
	}
}

func main() {
	loaded, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Не удалось загрузить задачи, начинаю с пустого списка: ", err)
		tasks = []task.Task{}
	} else {
		tasks = loaded
	}
	ReadInp()
}
