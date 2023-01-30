package database

import (
	"os"
	"sync"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)
var lock = &sync.Mutex{}

type dbInfo struct {

}

type DBConnection struct {
	DB *sql.DB
}

var dbInstance *DBConnection

func NewDBConnection() *DBConnection {
	if dbInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		db, err := sql.Open("mysql", getDBInfo().FormatDSN())

		if err == nil {
			dbInstance = &DBConnection{
				DB: db,
			}
		}

	}

	return dbInstance
}

func getDBInfo() *mysql.Config {
	return &mysql.Config {
		User: os.Getenv("DB_USER"),
        Passwd: os.Getenv("DB_PASS"),
        Net:    "tcp",
        Addr:   os.Getenv("DB_HOST"),
        DBName: os.Getenv("DB_NAME"),
	}
}