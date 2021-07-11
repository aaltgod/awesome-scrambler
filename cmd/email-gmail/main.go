package main

import (
	ag "github.com/alyaskastorm/awesome-scrambler/internal/emai-gmail"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Email-Gmail service is running")

	ag.RunApp()

	log.Println("Awesome-Scrambler service is shutdown")
}