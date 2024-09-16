package todo_test

import (
	"os"
	"path/filepath"
	"testing"

	todo "github.com/sebastianhevia/todo-cli"
)

// Helper function to create a new List with a temporary database for each test
func newTestList(t *testing.T) todo.List {
	t.Helper()
	dbName := "todo-test.db"

	l, err := todo.NewList(dbName)
	if err != nil {
		t.Fatalf("Failed to create new List: %v", err)
	}

	t.Cleanup(func() {
		os.RemoveAll(tempDir)
	})

	return l
}

func TestAdd(t *testing.T) {
	l := newTestList(t)

	taskName := "New Task"
	l.Add(taskName)

	task, err := l.Get(0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if task.Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, task.Task)
	}
}

func TestComplete(t *testing.T) {
	l := newTestList(t)

	taskName := "New Task"
	l.Add(taskName)

	task, err := l.Get(0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if task.Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, task.Task)
	}

	if task.Done {
		t.Errorf("New task should not be completed.")
	}
	l.Complete(0)

	if !task.Done {
		t.Errorf("New task should be completed.")
	}
}

func TestUpdate(t *testing.T) {
	l := newTestList(t)
	taskName := "New Task"
	l.Add(taskName)
	l.Update(0, "Updated Task")
	task, err := l.Get(0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if task.Task != "Updated Task" {
		t.Errorf("Expected %q, got %q instead", "Updated Task", task.Task)
	}
}

func TestList(t *testing.T) {
	l := newTestList(t)
	taskName := "New Task"
	l.Add(taskName)
	task, err := l.Get(0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if task.Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, task.Task)
	}
}
