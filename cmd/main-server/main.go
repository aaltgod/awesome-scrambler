package main

import (
	"log"

	"github.com/joho/godotenv"

	scrambler "github.com/aaltgod/awesome-scrambler/internal/main-server"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Main-server service is running")

	scrambler.RunApp()

	log.Println("Main-server service is shutdown")
}
