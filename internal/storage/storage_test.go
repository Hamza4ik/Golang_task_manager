package storage

import (
	"path/filepath"
	"testing"

	"github.com/Hamza4ik/Golang_task_manager/internal/task"
)

func TestSaveAndLoadTasks(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "tasks.json")

	tasks := []task.Task{
		{ID: 1, Text: "Test 1", Done: false},
		{ID: 2, Text: "Test 2", Done: true},
	}

	err := SaveTasks(tasks, path)
	if err != nil {
		t.Fatalf("SaveTasks failed: %v", err)
	}

	loaded, err := LoadTasks(path)
	if err != nil {
		t.Fatalf("Loads failed %v", err)
	}
	if len(loaded) != len(tasks) {
		t.Errorf("expected %d tasks, got %d", len(tasks), len(loaded))
	}
	if loaded[0].Text != tasks[0].Text {
		t.Errorf("expected first task text %q, got %q", tasks[0].Text, loaded[0].Text)
	}
}
