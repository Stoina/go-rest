package restservice

import (
	"log"
	"net/http"
	"strconv"

	repo "github.com/Stoina/go-rest-repository"
)

// Service exported
// Service ...
type Service struct {
	Name 			string
	Port 			int
	Repositories 	[]repo.Repository
}

// SecureService exported
// ...
type SecureService struct {
	RestService 	*Service
	CertFile 		string
	KeyFile 		string
}

// NewService exported
// ...
func NewService(name string, port int, repos []repo.Repository) *Service {
	return &Service {
		Name: name,
		Port: port,
		Repositories: repos}
}

// NewSecureService exported
// ...
func NewSecureService(name string, port int, repos []repo.Repository, certFile string, keyFile string) *SecureService {
	return &SecureService{
		RestService: &Service {
			Name: name,
			Port: port,
			Repositories: repos},
		CertFile: certFile,
		KeyFile: keyFile}
}

// Start exported
// Start ....
func (service *Service) Start() {
	service.initializeService()
	
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(service.Port), nil))
}

func (service *Service) initializeService() {

	log.Println("HTTP-Server port: " + strconv.Itoa(service.Port))

	for _, repo := range service.Repositories {
		log.Println("Handle Repository: " + repo.Name() + " with URL: " + repo.URL())

		controller := NewController(repo)

		http.Handle("/"+repo.URL()+"/", controller)
	}
}

// Start exported
// ...
func (secureService *SecureService) Start() {
	
	secureService.RestService.initializeService()

	log.Fatal(http.ListenAndServeTLS(":" + strconv.Itoa(secureService.RestService.Port), secureService.CertFile, secureService.KeyFile, nil))
}