package rest

import (
	"log"
	"net/http"
	"strconv"

	db "github.com/Stoina/go-database"
	controller "github.com/Stoina/go-rest/controller"
	repo "github.com/Stoina/go-rest/repository"
	config "github.com/Stoina/go-rest/config"
	testRepo "github.com/Stoina/go-rest/repository/test"
	testRepoConfig "github.com/Stoina/go-rest/repository/test/config"
	sqlRepo "github.com/Stoina/go-rest/repository/sql"
	sqlRepoConfig "github.com/Stoina/go-rest/repository/sql/config"
)

// Server exported
// Server ...
type Server struct {
	Port int

	Repositories []repo.Repository
}

// Start exported
// Start ....
func (server *Server) Start() {

	log.Println("Server port: " + strconv.Itoa(server.Port))

	for _, repo := range server.Repositories {
		log.Println("Handle Repository: " + repo.Name() + " with URL: " + repo.URL())

		controller := controller.NewController(repo)
		http.Handle("/"+repo.URL()+"/", controller)
	}

	portString := ":" + strconv.Itoa(server.Port)
	log.Fatal(http.ListenAndServe(portString, nil))
}

// NewServer exported
// NewServer ...
func NewServer(port int, repos []repo.Repository) *Server {
	return &Server {
		Port: port,
		Repositories: repos}
}

// NewServerFromConfigFile exported
// NewServerFromConfigFile ...
func NewServerFromConfigFile(configFilePath string) (*Server, error) {
	currentConfig, err := config.ReadConfigFromFile("json", configFilePath)

	if err != nil {
		return nil, err
	}

	convertedPort, err := strconv.Atoi(currentConfig.Server.Port)
	
	if err != nil {
		return nil, err
	}
	
	createdServer := &Server {
		Port: convertedPort,
		Repositories: createRepositoriesFromConfig(currentConfig)}

	return createdServer, nil
}

func createRepositoriesFromConfig(c *config.Config) []repo.Repository {
	repositoriesCount := len(c.TestRepositories) + len(c.SQLRepositories)
	repositories := make([]repo.Repository, repositoriesCount)

	repoCount := 0

	for _, testRepoConf := range c.TestRepositories {
		repositories[repoCount] = createTestRepositoryFromConfig(testRepoConf)
		repoCount++
	}

	for _, sqlRepoConf := range c.SQLRepositories {
		repositories[repoCount] = createSQLRepositoryFromConfig(sqlRepoConf)
		repoCount++
	}

	return repositories
}

func createTestRepositoryFromConfig(testRepoConf testRepoConfig.TestRepositoryConfig) *testRepo.TestRepository {
	return testRepo.NewTestRepository(testRepoConf.Name, testRepoConf.URL)
}

func createSQLRepositoryFromConfig(sqlRepoConf sqlRepoConfig.SQLRepositoryConfig) *sqlRepo.SQLRepository {

	dbConn, err := db.OpenDBConnection("postgres", "localhost", 5432, "postgres", "steinerj", "postgres")

	if err != nil {
		log.Println(err.Error())
	}

	return sqlRepo.NewSQLRepository(sqlRepoConf.Name, sqlRepoConf.URL, dbConn, &sqlRepoConf)
}