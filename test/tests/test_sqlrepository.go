package test


import (
	"log"
	
	rest "github.com/Stoina/go-rest"
	db "github.com/Stoina/go-database"
	repo "github.com/Stoina/go-rest/repository"
	sqlRepo "github.com/Stoina/go-rest/repository/sql"
	sqlRepoConfig "github.com/Stoina/go-rest/repository/sql/config"
)

func TestServerWithSQLRepository() {
	dbConn, err := db.OpenDBConnection("postgres", "localhost", 5432, "postgres", "steinerj", "postgres")

	if err != nil {
		log.Println(err.Error())
	}

	restServer := rest.Server{
		Port: 8080,

		Repositories: []repo.Repository{
			sqlRepo.NewSQLRepository("Golang-Test", "golangtest", dbConn, &sqlRepoConfig.SQLRepositoryConfig{Statement: "select * from golangtest"})}}

	restServer.Start()
}