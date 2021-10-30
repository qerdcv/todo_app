package main

import (
	"log"
	"todo_app/server"
)

func main() {
	s, err := server.New()
	if err != nil {
		log.Fatal(err)
	}
	s.Run()
}
