package restserver

import (
	"log"
	"testing"

	db "github.com/Stoina/go-database"
	repo "github.com/Stoina/go-rest-server/repo"
	sqlRepo "github.com/Stoina/go-rest-server/repo/sql"
)

func TestRestServerSQLRepository(t *testing.T) {
	dbConn, err := db.OpenDBConnection("postgres", "localhost", 5432, "postgres", "steinerj", "postgres")

	if err != nil {
		log.Println(err.Error())
	}

	restServer := NewServer("TestServer", 8080, []repo.Repository{getSQLRepository(dbConn)})
	restServer.Start()
}

func getSQLRepository(dbConn *db.Connection) *sqlRepo.SQLRepository {
	return sqlRepo.NewSQLRepository("Golang-Test", "golangtest", dbConn, sqlRepo.NewSQLRepositorySettings("Customer", nil))
}