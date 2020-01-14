package restserver

import (
	"io/ioutil"
	"log"
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
	} else if method == POST {
		handlePostRequest(rc, writer, request)
	}
}

func handleGetRequest(rc *Controller, writer http.ResponseWriter, request *http.Request) {
	header := writer.Header()
	header.Add("content-type", "application/json")
	writer.Write([]byte(rc.repo.Get(request.URL).ConvertToJSON()))
}

func handlePostRequest(rc *Controller, writer http.ResponseWriter, request *http.Request) {
	
	content, err := getContentFromRequest(request)

	if err != nil {
		log.Fatal(err)
	}

	contentType := getContentTypeFromRequest(request)
	
	header := writer.Header()
	header.Add("content-type", "application/json")
	writer.Write([]byte(rc.repo.Post(contentType, content).ConvertToJSON()))
}

func getContentFromRequest(request *http.Request) (string, error) {
	contentBytes, err := ioutil.ReadAll(request.Body)

	if err != nil {
		return "", err
	}
	
	return string(contentBytes), nil
}

func getContentTypeFromRequest(request *http.Request) string {
	contentType := request.Header.Get("Content-Type")
	return contentType
}