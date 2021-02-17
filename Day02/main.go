package main

import (
	"SofwareGoDay1/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	env := godotenv.Load("env.env")
	if env != nil {
		log.Fatal("Error loading .env file")
	}
	i := server.NewServer()
	i.Run()
}
