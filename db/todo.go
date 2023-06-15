package db

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	ID    int
	Title string
	Time  int
}

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository() (*TodoRepository, error) {
	db, err := sql.Open("sqlite3", "./data/dataStore.sqlite")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Create the TODO table if it doesn't exist
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS todo (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			time INT
		)
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create TODO table: %v", err)
	}

	return &TodoRepository{
		db: db,
	}, nil
}

func (r *TodoRepository) Close() error {
	return r.db.Close()
}

func (r *TodoRepository) Create(todo *Todo) error {
	insertQuery := "INSERT INTO todo (title,time) VALUES (?,?)"
	result, err := r.db.Exec(insertQuery, todo.Title, todo.Time)
	if err != nil {
		return fmt.Errorf("failed to insert TODO: %v", err)
	}

	id, _ := result.LastInsertId()
	todo.ID = int(id)

	return nil
}

func (r *TodoRepository) Read() ([]byte, error) {
	selectQuery := "SELECT id, title FROM todo"
	rows, err := r.db.Query(selectQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch TODOs: %v", err)
	}
	defer rows.Close()

	todos := make([]*Todo, 0)
	for rows.Next() {
		todo := &Todo{}
		err := rows.Scan(&todo.ID, &todo.Title)
		if err != nil {
			return nil, fmt.Errorf("failed to scan TODO: %v", err)
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate TODOs: %v", err)
	}

	// Convert todos slice to JSON
	jsonData, err := json.Marshal(todos)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal TODOs to JSON: %v", err)
	}

	return jsonData, nil
}


func (r *TodoRepository) Update(todo *Todo) error {
	updateQuery := "UPDATE todo SET title = ? WHERE id = ?"
	_, err := r.db.Exec(updateQuery, todo.Title, todo.ID)
	if err != nil {
		return fmt.Errorf("failed to update TODO: %v", err)
	}

	return nil
}

func (r *TodoRepository) Delete(id int) error {
	deleteQuery := "DELETE FROM todo WHERE id = ?"
	_, err := r.db.Exec(deleteQuery, id)
	if err != nil {
		return fmt.Errorf("failed to delete TODO: %v", err)
	}

	return nil
}