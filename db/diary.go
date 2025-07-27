package db

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/sainikmandal/diary-in-terminal/encryption"
)

func InitDB(db *sql.DB) error {
	createUsers := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		password_hash TEXT NOT NULL
	);`

	createEntries := `
	CREATE TABLE IF NOT EXISTS entries (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at TEXT NOT NULL,
		encrypted_text BLOB NOT NULL
	);`

	_, err := db.Exec(createUsers)
	if err != nil {
		return err
	}

	_, err = db.Exec(createEntries)
	return err
}

func InsertEntry(db *sql.DB, key []byte, content string) error {
	encData, err := encryption.Encrypt(content, key)
	if err != nil {
		return err
	}

	timestamp := time.Now().Format(time.RFC3339)
	_, err = db.Exec("INSERT INTO entries (created_at, encrypted_text) VALUES (?, ?)", timestamp, encData)
	return err
}

func ViewEntries(db *sql.DB, key []byte) error {
	rows, err := db.Query("SELECT id, created_at, encrypted_text FROM entries ORDER BY id DESC")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var createdAt string
		var encText []byte
		if err := rows.Scan(&id, &createdAt, &encText); err != nil {
			return err
		}

		decrypted, err := encryption.Decrypt(encText, key)
		if err != nil {
			decrypted = "[decryption failed]"
		}

		fmt.Printf("[%d] %s\n%s\n\n", id, createdAt, decrypted)
	}

	return nil
}

func DeleteEntry(db *sql.DB, id string) error {
	res, err := db.Exec("DELETE FROM entries WHERE id = ?", id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("no entry with that ID")
	}
	return nil
}
