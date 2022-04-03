package main

import (
	"log"

	"github.com/zloyboy/mongo/internal/server"
)

func main() {
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
