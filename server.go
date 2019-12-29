package rest

import (
	"log"
	"net/http"
	"strconv"

	controller "github.com/Stoina/go-rest/controller"
	repo "github.com/Stoina/go-rest/repository"
)

// Server exported
// Server ...
type Server struct {
	URL  string
	Port int
}

// Start exported
// Start ....
func (server *Server) Start(repo repo.Repository) {

	log.Println("Server url: " + server.URL)
	log.Println("Server port: " + strconv.Itoa(server.Port))

	log.Println("Repository: " + repo.Name())

	controller := controller.NewController(repo)

	http.Handle("/"+server.URL+"/", controller)

	portString := ":" + strconv.Itoa(server.Port)
	log.Fatal(http.ListenAndServe(portString, nil))
}
