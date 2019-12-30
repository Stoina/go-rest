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
