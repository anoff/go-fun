package main

import (
	"log"
	"os"
)

func main() {
	inputMessage := os.Getenv("INPUT")

	if inputMessage != "" {
		log.Println(inputMessage)
	}
}