package app

import (
	"bayudsatriyo/belajar-golang-restfull-api/helper"
	"database/sql"
	"time"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/golang-restful-api")
	helper.PanicIfError(err)
	db.SetMaxIdleConns(5)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(60 * time.Second)
	db.SetConnMaxIdleTime(10 * time.Second)
	return db
}
