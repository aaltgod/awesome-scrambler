package main

import (
	"log"

	scrambler "github.com/alyaskastorm/awesome-scrambler/internal/awesome-scrambler"
)

func main() {
	log.Println("Service is running")

	scrambler.RunApp()
}