package main

import "github.com/learning/api/server"

func main() {
	apiURL := "localhost:8080"
	server.StartAPI(apiURL)
}
