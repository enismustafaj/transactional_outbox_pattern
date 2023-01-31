package database

import (
	"os"
	"sync"
	"database/sql"
	"github.com/go-sql-driver/mysql"

	"github.com/transactional_outbox_pattern/main_service/model"
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

func (db DBConnection) insertData(user *model.User) {
	var connection *sql.DB = db.DB
	// TODO: ADD transaction

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