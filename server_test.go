package restserver

import (
	"log"
	"testing"

	db "github.com/Stoina/go-database"
	repo "github.com/Stoina/go-rest-server/repo"
)

func TestRestServerSQLRepository(t *testing.T) {
	dbConn, err := db.OpenDBConnection("postgres", "localhost", 5432, "postgres", "steinerj", "postgres")

	if err != nil {
		log.Println(err.Error())
	}

	restServer := NewServer("TestServer", 8080, []repo.Repository{repo.NewSQLRepository("Golang-Test", "golangtest", dbConn)})
	restServer.Start()
}
