package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Hamza4ik/Golang_task_manager/internal/storage"
	"github.com/Hamza4ik/Golang_task_manager/internal/task"
)

var tasks []task.Task

const tasksFile = "tasks.json"

func ReadInp() {
	fmt.Println("Write command: add,  list, mark, delete, edit, exit")
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
			if err := storage.SaveTasks(tasks, tasksFile); err != nil {
				fmt.Println("Не удалось сохранить")
			}
		case "mark":
			if len(parts) < 2 {
				fmt.Println("Len less then 2")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Id must be int")
				continue
			}
			task.MarkDone(&tasks, id)
			if err := storage.SaveTasks(tasks, tasksFile); err != nil {
				fmt.Println("Не удалось сохранить")
			}
		case "delete":
			if len(parts) < 2 {
				fmt.Println("Len less then 2")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Id must be int")
				continue
			}
			task.DeleteTask(&tasks, id)
			if err := storage.SaveTasks(tasks, tasksFile); err != nil {
				fmt.Println("Failed to save")
			}
		case "edit":
			partsEdit := strings.SplitN(line, " ", 3)
			if len(partsEdit) < 3 {
				fmt.Println("Len edit less then 3")
				continue
			}
			id, err := strconv.Atoi(partsEdit[1])
			if err != nil {
				fmt.Println("id must be int")
				continue
			}
			textEdit := partsEdit[2]
			task.EditTask(&tasks, id, textEdit)
			if err := storage.SaveTasks(tasks, tasksFile); err != nil {
				fmt.Println("Failed to save")
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
	loaded, err := storage.LoadTasks(tasksFile)
	if err != nil {
		fmt.Println("Не удалось загрузить задачи, начинаю с пустого списка: ", err)
		tasks = []task.Task{}
	} else {
		tasks = loaded
	}
	ReadInp()
}
