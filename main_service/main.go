package main

import (
	"log"
	"os"
	"net/http"
	"encoding/json"
	"github.com/joho/godotenv"

	"github.com/transactional_outbox_pattern/main_service/database"
	"github.com/transactional_outbox_pattern/main_service/model"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading env variables")
	}

}

var db *database.DBConnection

func main() {
	var serverPort string = os.Getenv("PORT")

	mux := http.NewServeMux()
	handler := http.HandlerFunc(handleRequest)

	mux.Handle("/api/data", handler)
	log.Println("Listening on Port: ", serverPort)

	db = database.NewDBConnection()
	http.ListenAndServe(serverPort, mux)

	if err := http.ListenAndServe(serverPort, mux); err != nil {
        log.Fatal(err)
    }
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var user model.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		db.InsertData(&user)
	} else {
		log.Println("Other request are not rally allowed")
	}
}