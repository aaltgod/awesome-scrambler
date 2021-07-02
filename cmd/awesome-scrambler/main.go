package main

import (
	scrambler "github.com/alyaskastorm/awesome-scrambler/internal/awesome-scrambler"
	"log"
)

func main() {
	log.Println("Service is running")

	scrambler.RunApp()
}