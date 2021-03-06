package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var (
	dbName = "health"
	dbUser = "root"
	dbPass = ""
)

func InitDb() (*sql.DB, error) {
	if envDbName := os.Getenv("DB_NAME"); envDbName != "" {
		dbName = envDbName
	}
	if envDbUser := os.Getenv("DB_USER"); envDbUser != "" {
		dbUser = envDbUser
	}
	if envDbPass := os.Getenv("DB_PASS"); envDbPass != "" {
		dbPass = envDbPass
	}
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

func GetBloodReadings(db *sql.DB, userId int64) ([]*BloodReading, error) {
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

func GetBloodReading(db *sql.DB, userId int64, readingId int64) (*BloodReading, error) {
	sql := `
		SELECT id, user_id, systolic, diastolic, pulse, reading_date
		FROM blood_pressures
		WHERE id = ? AND user_id = ?
	`
	row := db.QueryRow(sql, readingId, userId)
	var id int64
	var systolic int32
	var diastolic int32
	var pulse int32
	var readingDate string
	if err := row.Scan(&id, &userId, &systolic, &diastolic, &pulse, &readingDate); err != nil {
		log.Printf("err: %s", err)
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
	return &reading, nil
}

func CreateBloodReading(db *sql.DB, userId int64, systolic, diastolic, pulse int32, date string) error {
	sql := `
		INSERT INTO blood_pressures
		(id, user_id, systolic, diastolic, pulse, reading_date, created_at, updated_at, deleted_at) 
		VALUES(NULL, ?, ?, ?, ?, ?, NOW(), NOW(), NULL);
	`
	result, err := db.Exec(sql, userId, systolic, diastolic, pulse, date)
	if err != nil {
		return err
	}
	if _, err = result.RowsAffected(); err != nil {
		return err
	}
	return nil
}

func UpdateBloodReading(db *sql.DB, userId int64, readingId int64, systolic, diastolic, pulse int32, date string) error {
	sql := `
		UPDATE blood_pressures
		SET systolic = ?, diastolic = ?, pulse = ?, reading_date = ?, updated_at = NOW()
		WHERE user_id = ? AND id = ?
	`
	result, err := db.Exec(sql, systolic, diastolic, pulse, date, userId, readingId)
	if err != nil {
		return err
	}
	if _, err = result.RowsAffected(); err != nil {
		return err
	}
	return nil
}

func DeleteBloodReading(db *sql.DB, userId int64, readingId int64) error {
	sql := `
		UPDATE blood_pressures
		SET deleted_at = NOW()
		WHERE user_id = ? AND id = ?
	`
	result, err := db.Exec(sql, userId, readingId)
	if err != nil {
		return err
	}
	if _, err = result.RowsAffected(); err != nil {
		return err
	}
	return nil
}

func GetWeightEntries(db *sql.DB, userId int64) ([]*WeightEntry, error) {
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

func GetWeightEntry(db *sql.DB, userId int64, entryId int64) (*WeightEntry, error) {
	sql := `
		SELECT id, user_id, weight, entered_date
		FROM weights
		WHERE id = ? AND user_id = ?
	`
	row := db.QueryRow(sql, entryId, userId)
	var id int64
	var weight float32
	var enteredDate string
	if err := row.Scan(&id, &userId, &weight, &enteredDate); err != nil {
		log.Printf("err: %s", err)
		return nil, err
	}
	entry := WeightEntry{
		Id:     id,
		UserId: userId,
		Weight: weight,
		Date:   enteredDate,
	}
	return &entry, nil
}

func CreateWeightEntry(db *sql.DB, userId int64, weight float32, date string) error {
	sql := `
		INSERT INTO weights
		(id, user_id, weight, entered_date, deleted_at, created_at, updated_at) 
		VALUES(NULL, ?, ?, ?, NULL, NOW(), NOW());
	`
	result, err := db.Exec(sql, userId, weight, date)
	if err != nil {
		return err
	}
	if _, err = result.RowsAffected(); err != nil {
		return err
	}
	return nil
}

func UpdateWeightEntry(db *sql.DB, userId, entryId int64, weight float32, date string) error {
	sql := `
		UPDATE weights
		SET weight = ?, entered_date = ?, updated_at = NOW()
		WHERE user_id = ? AND id = ?
	`
	result, err := db.Exec(sql, weight, date, userId, entryId)
	if err != nil {
		return err
	}
	if _, err = result.RowsAffected(); err != nil {
		return err
	}
	return nil
}

func DeleteWeightEntry(db *sql.DB, userId, entryId int64) error {
	sql := `
		UPDATE weights
		SET deleted_at = NOW()
		WHERE user_id = ? AND id = ?
	`
	result, err := db.Exec(sql, userId, entryId)
	if err != nil {
		return err
	}
	if _, err = result.RowsAffected(); err != nil {
		return err
	}
	return nil
}

func RegisterUser(db *sql.DB, name, lastName, email, pass string) error {
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
