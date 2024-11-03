package main

import "github.com/vincentandreas/dealls/api/handlers"

func main() {
	router, _ := handlers.SetupServer()
	router.Run("localhost:8080")
}