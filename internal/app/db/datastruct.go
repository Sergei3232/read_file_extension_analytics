package db

import "time"

type FileDataBase struct {
	Id   int       `db:"id"`
	File string    `db:"path"`
	Date time.Time `db:"created_at"`
}
