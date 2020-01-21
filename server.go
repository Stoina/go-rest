package restserver

import (
	"log"
	"net/http"
	"strconv"

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