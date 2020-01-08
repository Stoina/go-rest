package restserver

import (
	"net/http"
	
	repo "github.com/Stoina/go-rest-server/repo"
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

	requestMethod := getRequestMethod(request.Method)

	if requestMethod != UNDEFINED {
		handleRequest(requestMethod, controller, writer, request)
	}
}

func getRequestMethod(restMethodToParse string) RequestMethod {
	return ParseToRequestMethod(restMethodToParse)
}

func handleRequest(method RequestMethod, rc *Controller, writer http.ResponseWriter, request *http.Request) {
	if method == GET {
		handleGetRequest(rc, writer, request)
	}
}

func handleGetRequest(rc *Controller, writer http.ResponseWriter, request *http.Request) {
	header := writer.Header()
	header.Add("content-type", "application/json")
	writer.Write([]byte(rc.repo.Get(request.URL).ConvertToJSON()))
}

func handlePutRequest() {

}
