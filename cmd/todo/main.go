package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	todo "github.com/sebastianhevia/todo-cli"
)

func main() {
	add := flag.String("add", "", "Add a new task")
	complete := flag.Int("complete", 0, "Mark a task as completed")
	list := flag.Bool("list", false, "List all tasks")
	delete := flag.Int("delete", 0, "Delete a task")
	update := flag.String("update", "", "Update a task")
	id := flag.Int("id", 0, "ID of the task to update")
	flag.Parse()

	dbPath := os.Getenv("TODO_DB_PATH")
	if dbPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to get user home directory:", err)
			os.Exit(1)
		}
		dbPath = filepath.Join(homeDir, ".todo", "todo.db")
	}

	l, err := todo.New(dbPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer l.Close()

	switch {
	case *add != "":
		err = l.Add(*add)
	case *complete > 0:
		err = l.Complete(*complete)
	case *list:
		items, err := l.List()
		if err == nil {
			for _, item := range items {
				fmt.Printf("%d: %s (Done: %t)\n", item.ID, item.Task, item.Done)
			}
		}
	case *delete > 0:
		err = l.Delete(*delete)
	case *update != "" && *id > 0:
		err = l.Update(*id, *update)
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
