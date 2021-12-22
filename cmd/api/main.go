package main

import "github.com/haitien/chi/server"

func main() {
	server := server.NewServer()
	server.Serve("8080")
}
