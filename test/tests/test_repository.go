package test

import (
	"log"

	rest "github.com/Stoina/go-rest"
	repo "github.com/Stoina/go-rest/repository"
	testRepo "github.com/Stoina/go-rest/repository/test"
)

func TestServerWithRestRepository() {

	log.Println("Testing rest server...")

	restServer := rest.Server{
		Port: 8080,

		Repositories: []repo.Repository{
			testRepo.NewTestRepository("Test-One", "Test-One"),
			testRepo.NewTestRepository("Test-Two", "Test-Two")}}

	restServer.Start()
}