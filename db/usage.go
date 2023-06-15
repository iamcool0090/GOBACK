package db

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Usage struct {
	SessionID int `json:"sessionID"`
	FromTime  int `json:"fromTime"`
	ToTime    int `json:"toTime"`
}

type UsageRepository struct {
	db *sql.DB
}

func NewUsageRepository() (*UsageRepository, error) {
	db, err := sql.Open("sqlite3", "./data/dataStore.sqlite")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Create the usage table if it doesn't exist
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS usage (
			sessionID INTEGER PRIMARY KEY AUTOINCREMENT,
			fromTime INT,
			toTime INT
		)
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create usage table: %v", err)
	}

	return &UsageRepository{
		db: db,
	}, nil
}

func (r *UsageRepository) Close() error {
	return r.db.Close()
}

func (r *UsageRepository) CreateUsage(usage *Usage) error {
	insertQuery := "INSERT INTO usage (fromTime) VALUES (?)"
	result, err := r.db.Exec(insertQuery, usage.FromTime)
	if err != nil {
		return fmt.Errorf("failed to insert usage: %v", err)
	}

	sessionID, _ := result.LastInsertId()
	usage.SessionID = int(sessionID)

	return nil
}

func (r *UsageRepository) ReadUsage() ([]byte, error) {
	selectQuery := "SELECT sessionID, fromTime, toTime FROM usage"
	rows, err := r.db.Query(selectQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch usage: %v", err)
	}
	defer rows.Close()

	usages := []Usage{}
	for rows.Next() {
		usage := Usage{}
		err := rows.Scan(&usage.SessionID, &usage.FromTime, &usage.ToTime)
		if err != nil {
			return nil, fmt.Errorf("failed to scan usage row: %v", err)
		}
		usages = append(usages, usage)
	}

	// Convert usage slice to JSON
	jsonData, err := json.Marshal(usages)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal usage to JSON: %v", err)
	}

	return jsonData, nil
}

func (r *UsageRepository) UpdateToTime(usage *Usage) error {
	updateQuery := "UPDATE usage SET toTime = ? WHERE sessionID = ?"
	_, err := r.db.Exec(updateQuery, usage.ToTime, usage.SessionID)
	if err != nil {
		return fmt.Errorf("failed to update usage: %v", err)
	}

	return nil
}

