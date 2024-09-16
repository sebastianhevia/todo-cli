package todo

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func getDBPath(dbName string) string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".todo", dbName)
}

func initDB(dbName string) (*sql.DB, error) {
	dbPath := getDBPath(dbName)
	os.MkdirAll(filepath.Dir(dbPath), os.ModePerm)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS todos (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            task TEXT NOT NULL,
            done BOOLEAN NOT NULL DEFAULT 0,
            created_at DATETIME NOT NULL,
            completed_at DATETIME
        )
    `)
	if err != nil {
		return nil, err
	}

	return db, nil
}
