package event_reader

import (
	"os"
	"log"
	"database/sql"
	"time"
	"github.com/go-sql-driver/mysql"
)

type Reader struct {
	DBConnection *sql.DB
	ErrorChan chan string
}

func NewReader() *Reader {
	db, err := sql.Open("mysql", getDBInfo().FormatDSN())

	if err != nil {
		log.Fatal("Error creating db connection")
	}

	reader := &Reader{
		DBConnection: db,
		ErrorChan: make(chan string),
	}

	return reader
}

func (r *Reader) Start() {
	ticker := time.NewTicker(1000 * time.Millisecond)

	for {
		select {
		case err := <- r.ErrorChan:
			log.Fatal(err)
			return
		case <-ticker.C:
			r.readEvents()
		}
	}
}

func (r *Reader) readEvents() {
	var connection *sql.DB = r.DBConnection

	_, err := connection.Query("select * from outbox_table limit 2")

	if err != nil {
		log.Fatal("Error fetching events from outbox table", err)
	}

	log.Println("Events fetched from outbox table")
	time.Sleep(3000 * time.Millisecond)
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