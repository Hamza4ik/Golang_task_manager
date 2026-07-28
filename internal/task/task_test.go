package task

import (
	"testing"
)

func TestAddTask(t *testing.T) {
	var tasks []Task
	AddTask(&tasks, Task{ID: 1, Text: "Test", Done: false})

	if len(tasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(tasks))
	}
}

func TestMarkDone(t *testing.T) {
	var tasks []Task
	AddTask(&tasks, Task{ID: 1, Text: "Test", Done: false})
	MarkDone(&tasks, 1)

	if !tasks[0].Done {
		t.Errorf("expected task 1 to be marked done")
	}
}

func TestDeleteTask(t *testing.T) {
	var tasks []Task
	AddTask(&tasks, Task{ID: 1, Text: "Test 1", Done: false})
	AddTask(&tasks, Task{ID: 2, Text: "Test 2", Done: false})
	DeleteTask(&tasks, 2)
	if len(tasks) != 1 {
		t.Errorf("expected 1 task after delete, got %d", len(tasks))
	}
}

func TestEditTask(t *testing.T) {
	var tasks []Task
	AddTask(&tasks, Task{ID: 1, Text: "Text", Done: false})
	EditTask(&tasks, 1, "New text")

	if tasks[0].Text != "New text" {
		t.Errorf("expected text to be updated, got %s", tasks[0].Text)
	}
}
