package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Starting the application...")

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Executing task at", time.Now())
		}
	}
}
