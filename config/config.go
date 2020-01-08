package restserver

import (
	"encoding/json"
	"errors"
	"os"
)

// Config exported
// Config ...
type Config struct {
	
	Server struct {
		Name string `json:"Name"`
		Port int `json:Port`
	} `json:"RestServer"`
	
	SQLRepositories []SQLRepositoryConfig `json:"SQLRepositories"`
}

// ReadConfigFromFile epxported
// ReadConfigFromFile ...
func ReadConfigFromFile(encoding string, filePath string) (*Config, error) {
	f, err := os.Open(filePath)
	
	if err != nil {
		return nil, err
	}

	if encoding == "json" {
		return readConfigFromJSONFile(f)
	}

	return nil, errors.New("could not read file encoding: " + encoding)
}

func readConfigFromJSONFile(file *os.File) (*Config, error) {

	defer file.Close()
	
	decoder := json.NewDecoder(file)

	var config Config
	err := decoder.Decode(&config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
