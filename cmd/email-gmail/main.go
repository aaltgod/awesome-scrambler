package main

import (
	"log"

	ag "github.com/aaltgod/awesome-scrambler/internal/email-gmail"
	"github.com/joho/godotenv"
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
