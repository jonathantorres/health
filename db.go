package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"log"
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

func getBloodReadings(db *sql.DB, userId int64) ([]*BloodReading, error) {
	sql := `
		SELECT id, user_id, systolic, diastolic, pulse, reading_date
		FROM blood_pressures
		WHERE user_id = ? AND deleted_at IS NULL
		ORDER BY reading_date DESC;
	`
	rows, err := db.Query(sql, userId)
	if err != nil {
		log.Printf("err: %s", err)
		return nil, err
	}
	defer rows.Close()
	readings := make([]*BloodReading, 0)
	for rows.Next() {
		var id int64
		var userId int64
		var systolic int32
		var diastolic int32
		var pulse int32
		var readingDate string
		if err := rows.Scan(&id, &userId, &systolic, &diastolic, &pulse, &readingDate); err != nil {
			log.Printf("%s", err)
			return nil, err
		}
		reading := BloodReading{
			Id:        id,
			UserId:    userId,
			Systolic:  systolic,
			Diastolic: diastolic,
			Pulse:     pulse,
			Date:      readingDate,
		}
		readings = append(readings, &reading)
	}
	return readings, nil
}

func getBloodReading(db *sql.DB, userId int64, readingId int64) (*BloodReading, error) {
	// todo
	return nil, nil
}

func getWeightEntries(db *sql.DB, userId int64) ([]*WeightEntry, error) {
	sql := `
		SELECT id, user_id, weight, entered_date
		FROM weights
		WHERE user_id = ? AND deleted_at IS NULL
		ORDER BY entered_date DESC;
	`
	rows, err := db.Query(sql, userId)
	if err != nil {
		log.Printf("err: %s", err)
		return nil, err
	}
	defer rows.Close()
	entries := make([]*WeightEntry, 0)
	for rows.Next() {
		var id int64
		var userId int64
		var weight float32
		var enteredDate string
		if err := rows.Scan(&id, &userId, &weight, &enteredDate); err != nil {
			log.Printf("%s", err)
			return nil, err
		}
		entry := WeightEntry{
			Id:     id,
			UserId: userId,
			Weight: weight,
			Date:   enteredDate,
		}
		entries = append(entries, &entry)
	}
	return entries, nil
}

func getWeightEntry(db *sql.DB, userId int64, entryId int64) (*WeightEntry, error) {
	// todo
	return nil, nil
}

func registerUser(db *sql.DB, name, lastName, email, pass string) error {
	sql := `
		INSERT INTO users
		(id, name, last_name, email, password, remember_token, created_at, updated_at, deleted_at) 
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
