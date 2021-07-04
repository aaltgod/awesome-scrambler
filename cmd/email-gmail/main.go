package main

import (
	ag "github.com/alyaskastorm/awesome-scrambler/internal/emai-gmail"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalln(err)
	}

	ag.RunApp()
}