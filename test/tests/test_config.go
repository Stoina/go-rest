package test

import (
	"log"

	rest "github.com/Stoina/go-rest"
)

func TestServerWithConfiguration() {

	restServer, err := rest.NewServerFromConfigFile("C:\\Users\\steinerj\\go\\src\\github.com\\stoina\\go-rest\\test\\config.json")

	if err != nil {
		log.Println(err.Error())
	}
	
	restServer.Start()
}