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

func MarkDone(tasks *[]Task, id int) {
	found := false
	for i := range *tasks {
		if (*tasks)[i].ID == id {
			(*tasks)[i].Done = true
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Id not found")
	}
}

func DeleteTask(tasks *[]Task, id int) {
	found := false
	for i := range *tasks {
		if (*tasks)[i].ID == id {
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Task not found")
	}
}

func EditTask(tasks *[]Task, id int, newText string) {
	found := false
	for i := range *tasks {
		if (*tasks)[i].ID == id {
			(*tasks)[i].Text = newText
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Task not found")
	}
}
