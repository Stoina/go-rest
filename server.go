package restserver

import (
	"log"
	"net/http"
	"strconv"

	db "github.com/Stoina/go-database"
	config "github.com/Stoina/go-rest-server/config"
	repo "github.com/Stoina/go-rest-server/repo"
)

// Server exported
// Server ...
type Server struct {
	Name string
	Port int
	Repositories []repo.Repository
}

// Start exported
// Start ....
func (server *Server) Start() {

	log.Println("Server port: " + strconv.Itoa(server.Port))

	for _, repo := range server.Repositories {
		log.Println("Handle Repository: " + repo.Name() + " with URL: " + repo.URL())

		controller := NewController(repo)
		http.Handle("/"+repo.URL()+"/", controller)
	}

	portString := ":" + strconv.Itoa(server.Port)
	log.Fatal(http.ListenAndServe(portString, nil))
}

// NewServer exported
// NewServer ...
func NewServer(name string, port int, repos []repo.Repository) *Server {
	return &Server {
		Name: name,
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
	
	createdServer := &Server {
		Name: currentConfig.Server.Name,
		Port: currentConfig.Server.Port,
		Repositories: createRepositoriesFromConfig(currentConfig)}

	return createdServer, nil
}

func createRepositoriesFromConfig(c *config.Config) []repo.Repository {
	repositories := make([]repo.Repository, len(c.SQLRepositories))

	repoCount := 0
	
	for _, sqlRepoConf := range c.SQLRepositories {
		repositories[repoCount] = createSQLRepositoryFromConfig(sqlRepoConf)
		repoCount++
	}

	return repositories
}

func createSQLRepositoryFromConfig(sqlRepoConf config.SQLRepositoryConfig) *repo.SQLRepository {

	dbConn, err := db.OpenDBConnection("postgres", "localhost", 5432, "postgres", "steinerj", "postgres")

	if err != nil {
		log.Println(err.Error())
	}

	return repo.NewSQLRepository(sqlRepoConf.Name, sqlRepoConf.URL, dbConn)
}