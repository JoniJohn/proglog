package main

import (
	"log"

	"github.com/jonijohn/proglog/internal/server"
)

func main() {
	srv := server.NewHttpServer(":8082")
	log.Fatal(srv.ListenAndServe())
}
