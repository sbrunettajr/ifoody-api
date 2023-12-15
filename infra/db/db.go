package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	config := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		DBName:               "ifoody_db",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
