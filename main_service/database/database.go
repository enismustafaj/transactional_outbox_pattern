package database

import (
	"os"
	"sync"
	"log"
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

		connectionError := db.Ping()

		if connectionError != nil {
			log.Fatal("Error connecting to db: ", connectionError)
		}

		if err == nil {
			dbInstance = &DBConnection{
				DB: db,
			}
		}

	}

	return dbInstance
}

func (db DBConnection) InsertData(user *model.User) {
	var connection *sql.DB = db.DB
	tx, err := connection.Begin()

	if err != nil {
		log.Fatal("Error creating transaction")
	}

	entityId, userDataInsertError := tx.Exec(`insert into user_table (FirstName, LastName) values (?, ?)`, user.FirstName, user.LastName)

	if userDataInsertError != nil  {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("unable to rollback: %v", rollbackErr)
		}
		log.Fatal("inserting user data error: ", userDataInsertError)
	}

	entityIdResult, _ := entityId.LastInsertId()
	_, eventDataInsertError := tx.Exec(`insert into outbox_table (Event, EntityId) values (?, ?)`, "event", entityIdResult)

	if eventDataInsertError != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("unable to rollback: %v", rollbackErr)
		}
		log.Fatal("inserting event data error: ", eventDataInsertError)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
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