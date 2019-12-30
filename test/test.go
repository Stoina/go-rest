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
		Port: 12345,

		Repositories: []repo.Repository{
			repo.NewTestRepository("Test-One", "Test-One"),
			repo.NewTestRepository("Test-Two", "Test-Two")}}

	restServer.Start()
}
