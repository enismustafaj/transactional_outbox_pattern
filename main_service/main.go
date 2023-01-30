package main

import (
	"log"
	"os"
	"net/http"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading env variables")
	}

	
}

func main() {
	var serverPort string = os.Getenv("PORT")

	mux := http.NewServeMux()
	handler := http.HandlerFunc(handleRequest)

	mux.Handle("/api/data", handler)

	http.ListenAndServe(serverPort, mux)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Println("post request")
	} else {
		log.Println("Other request")
	}
}