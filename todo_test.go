package todo_test

import (
    "io/ioutil"
    "os"
    "testing"

    todo "github.com/sebastianhevia/todo-cli"
)

func TestAdd(t *testing.T) {
    l := todo.List{}

    taskName := "New Task"
    l.Add(taskName)

    if l[0].Task != taskName {
        t.Errorf("Expected %q, got %q instead", taskName, l[0].Task)
    }
}


