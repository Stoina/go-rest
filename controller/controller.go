package rest

import (
	"net/http"
	"net/url"

	methods "github.com/Stoina/go-rest/methods"
	repo "github.com/Stoina/go-rest/repository"
)

// Controller exported
// Controller ....
type Controller struct {
	repo repo.Repository
}

// NewController exported
// NewController ...
func NewController(repo repo.Repository) *Controller {
	return &Controller{repo: repo}
}

// ServeHTTP exported
// ServeHTTP
func (controller *Controller) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	restMethod := getRequestMethod(request.Method)

	switch *restMethod {

	case methods.GET:
		handleGetCall(controller, writer, request.URL)

	case methods.PUT:
		handlePutCall()
	}

}

func getRequestMethod(restMethodToParse string) *methods.Method {
	return methods.ParseFromString(restMethodToParse)
}

func handleGetCall(rc *Controller, writer http.ResponseWriter, calledURL *url.URL) {

	header := writer.Header()
	header.Add("content-type", "application/json")
	writer.Write([]byte(rc.repo.Get(calledURL)))

}

func handlePutCall() {

}
