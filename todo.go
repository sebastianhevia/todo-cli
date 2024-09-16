package todo

import (
	"database/sql"
	"time"

	database "github.com/sebastianhevia/todo-cli/database"
)

type item struct {
	ID          int
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt sql.NullTime
}

type List struct {
	db *sql.DB
}

func NewList(dbName string) (*List, error) {
	db, err := database.initDB(dbName)
	if err != nil {
		return nil, err
	}
	return &List{db: db}, nil
}

func (l *List) Add(task string) error {
	_, err := l.db.Exec("INSERT INTO todos (task, created_at) VALUES (?, ?)", task, time.Now())
	return err
}

func (l *List) Complete(id int) error {
	_, err := l.db.Exec("UPDATE todos SET done = 1, completed_at = ? WHERE id = ?", time.Now(), id)
	return err
}

func (l *List) List() ([]item, error) {
	rows, err := l.db.Query("SELECT id, task, done, created_at, completed_at FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []item
	for rows.Next() {
		var i item
		err := rows.Scan(&i.ID, &i.Task, &i.Done, &i.CreatedAt, &i.CompletedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

func (l *List) Update(id int, task string) error {
	_, err := l.db.Exec("UPDATE todos SET task = ? WHERE id = ?", task, id)
	return err
}

func (l *List) Delete(id int) error {
	_, err := l.db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}

func (l *List) Close() error {
	return l.db.Close()
}

func (l *List) Get(id int) (item, error) {
	row := l.db.QueryRow("SELECT id, task, done, created_at, completed_at FROM todos WHERE id = ?", id)
	var i item
	err := row.Scan(&i.ID, &i.Task, &i.Done, &i.CreatedAt, &i.CompletedAt)
	return i, err
}
