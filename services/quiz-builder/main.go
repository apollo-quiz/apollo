package main

import (
	"log"
	"quiz-builder/web"
)

func main() {
	err := web.StartServer()
	if err != nil {
		log.Println("Failed to start")
		return
	}
}
