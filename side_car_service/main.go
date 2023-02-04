package main

import (
	"github.com/joho/godotenv"

	"github.com/transactional_outbox_pattern/side_car_service/event_reader"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading env variables")
	}

}

func main() {
	reader := event_reader.NewReader()
	reader.Start()
}