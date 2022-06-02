package main

import (
	"github.com/decadev/squad10/healthplus/db"
	"github.com/decadev/squad10/healthplus/router"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Something went wrong")
	}

	db.SetupDB()

	router.SetupRouter()
}
