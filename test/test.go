package main

import (
	"log"

	rest "github.com/Stoina/go-rest"
	repo "github.com/Stoina/go-rest/repository"
)

func main() {
	testRestRepository()
}

func testRestRepository() {

	log.Println("Test rest server")

	restServer := rest.Server{
		URL:  "rest-test-server",
		Port: 12345}

	serverRepo := repo.TestRepository{}

	restServer.Start(serverRepo)
}