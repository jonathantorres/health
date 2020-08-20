package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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
