package main

import (
	"fmt"
	"log"

	db "github.com/Stoina/go-database"
	rest "github.com/Stoina/go-rest"
	repo "github.com/Stoina/go-rest/repository"
	sqlRepo "github.com/Stoina/go-rest/repository/sql"
	testRepo "github.com/Stoina/go-rest/repository/test"
)

func main() {
	testSQLRepository()
}

func testRestRepository() {

	log.Println("Testing rest server...")

	restServer := rest.Server{
		Port: 8080,

		Repositories: []repo.Repository{
			testRepo.NewTestRepository("Test-One", "Test-One"),
			testRepo.NewTestRepository("Test-Two", "Test-Two")}}

	restServer.Start()
}

func testSQLRepository() {
	dbConn, err := db.OpenDBConnection("postgres", "localhost", 5432, "postgres", "steinerj", "postgres")

	if err != nil {
		fmt.Println(err.Error())
	}

	restServer := rest.Server{
		Port: 8080,

		Repositories: []repo.Repository{
			sqlRepo.NewSQLRepository("Golang-Test", "golangtest", dbConn, &sqlRepo.SQLRepositoryConfig{Statement: "select * from golangtest"})}}

	restServer.Start()
}
