package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	inputMessage := os.Getenv("inputMessage")
	outputMessageLocation := os.Getenv("outputQueueItem")

	if inputMessage != "" {
		// just write the length to the queue location. Just to prove this works.
		cacheFile, err := os.OpenFile(outputMessageLocation, os.O_WRONLY|os.O_CREATE, 0)
		if err != nil {
			log.Fatalf("Unable to open file %s %s", outputMessageLocation, err)
		}
		l := len(inputMessage)
		s := strconv.Itoa(l)
		cacheFile.WriteString(s)
	}
}