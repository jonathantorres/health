package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	dbName = "health"
	dbUser = "root"
	dbPass = ""
)

func initDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func registerUser(db *sql.DB, name, lastName, email, pass string) error {
	sql := `
		INSERT INTO users
		(id, NAME, last_name, email, PASSWORD, remember_token, created_at, updated_at, deleted_at) 
		VALUES(NULL, ?, ?, ?, ?, NULL, NOW(), NOW(), NULL);
	`
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	result, err := db.Exec(sql, name, lastName, email, hashedPass)
	if err != nil {
		return err
	}
	if _, err = result.RowsAffected(); err != nil {
		return err
	}
	return nil
}
