package main

import (
	"github.com/joho/godotenv"
	"log"

	scrambler "github.com/alyaskastorm/awesome-scrambler/internal/main-server"
)

func main() {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Awesome-Scrambler service is running")

	scrambler.RunApp()

	log.Println("Awesome-Scrambler service is shutdown")
}