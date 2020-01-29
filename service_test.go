package restservice

import (
	"log"
	"testing"

	db "github.com/Stoina/go-database"
	repo "github.com/Stoina/go-rest-repository"
)

func TestRestServerSQLRepository(t *testing.T) {
	dbConn, err := db.OpenDBConnection("postgres", "localhost", 5432, "postgres", "steinerj", "postgres")

	if err != nil {
		log.Println(err.Error())
	}

	restServer := NewService("TestServer", 8080, []repo.Repository{getSQLRepository(dbConn)})
	restServer.Start()
}

func getSQLRepository(dbConn *db.Connection) *repo.SQLRepository {
	return repo.NewSQLRepository("Golang-Test", "golangtest", dbConn, repo.NewSQLRepositorySetting("Customer", nil))
}