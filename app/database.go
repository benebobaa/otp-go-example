package app

import (
	"database/sql"
	"sent-email-otp/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://root:root@localhost:5432/email_otp?sslmode=disable")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
