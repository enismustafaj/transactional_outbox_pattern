package main

import (
	"log"
	"os"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"

	"github.com/transactional_outbox_pattern/main_service/database"
	"github.com/transactional_outbox_pattern/main_service/handlers"
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
	server := gin.Default()
	server.POST("/api/data", handlers.)

	log.Println("Listening on Port: ", serverPort)

	db = database.NewDBConnection()

	server.Run("localhost:" + serverPort)

	defer db.DB.Close()
}