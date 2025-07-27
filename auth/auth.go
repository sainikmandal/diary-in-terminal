package auth

import (
	"crypto/sha256"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func Register(db *sql.DB, password string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("INSERT INTO users (password_hash) VALUES (?)", string(hashed))
	if err != nil {
		return nil, err
	}

	return keyFromPassword(password), nil
}

func Authenticate(db *sql.DB, password string) ([]byte, error) {
	row := db.QueryRow("SELECT password_hash FROM users ORDER BY id DESC LIMIT 1")

	var storedHash string
	err := row.Scan(&storedHash)
	if err != nil {
		return nil, errors.New("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		return nil, errors.New("incorrect password")
	}

	return keyFromPassword(password), nil
}

func keyFromPassword(pw string) []byte {
	hash := sha256.Sum256([]byte(pw))
	return hash[:]
}
