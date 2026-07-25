package task

import "fmt"

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func AddTask(tasks *[]Task, t Task) {
	*tasks = append(*tasks, t)
}

func ListTasks(tasks []Task) {
	for _, t := range tasks {
		fmt.Printf("[%d] %s - done: %t\n", t.ID, t.Text, t.Done)
	}
}
