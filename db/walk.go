package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type File struct {
	Name    string
	Path    string
	ModTime int64
	Size    int64
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS files (
		name TEXT,
		path TEXT,
		mod_time INTEGER,
		size INTEGER
	)`)
	return err
}

func insertFile(db *sql.DB, file *File) error {
	stmt, err := db.Prepare("INSERT INTO files(name, path, mod_time, size) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(file.Name, file.Path, file.ModTime, file.Size)
	return err
}

func getFile(db *sql.DB, name string) (*File, error) {
	row := db.QueryRow("SELECT name, path, mod_time, size FROM files WHERE name=?", name)

	var file File
	err := row.Scan(&file.Name, &file.Path, &file.ModTime, &file.Size)
	if err != nil {
		return nil, err
	}

	return &file, nil
}

func updateFile(db *sql.DB, name string, file *File) error {
	stmt, err := db.Prepare("UPDATE files SET name=?, path=?, mod_time=?, size=? WHERE name=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(file.Name, file.Path, file.ModTime, file.Size, name)
	return err
}

func deleteFile(db *sql.DB, name string) error {
	stmt, err := db.Prepare("DELETE FROM files WHERE name=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name)
	return err
}