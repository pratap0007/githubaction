package main

import (
	"log"
	"net"
	"net/http"

	"github.com/pratap0007/githubaction/dashboard"
)

func main() {

	log.Println("starting the server")

	dashboard.Test()

	http.HandleFunc("/dashboard", dashboard.Details)
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	http.Serve(listener, nil)
}
