package main

import (
	"log"

	scrambler "github.com/alyaskastorm/awesome-scrambler/internal/awesome-scrambler"
)

func main() {
	log.Println("Awesome-Scrambler service is running")

	scrambler.RunApp()

	log.Println("Awesome-Scrambler service is shutdown")
}