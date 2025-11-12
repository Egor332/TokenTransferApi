package main

import (
	"log"

	"github.com/Egor332/TokenTransferApi/database"
)

func main() {
	database.Connect()

	sqlDB, err := database.DB.DB()
	if err != nil {

	}
	defer sqlDB.Close()

	log.Println("Database connection is ready")

	log.Println("Server start")
}
