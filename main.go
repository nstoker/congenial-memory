package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found %s", err)
	}
}

func main() {
	log.Println("ok, compiled cleanly")
}
