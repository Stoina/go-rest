package restconfig

import (
	"encoding/json"
	"errors"
	"os"
	
	testRepoConfig "github.com/Stoina/go-rest/repository/test/config"
	sqlRepoConfig "github.com/Stoina/go-rest/repository/sql/config"
)

// Config exported
// Config ...
type Config struct {
	
	Server struct {
		Port string `json:"Port"`
	} `json:"RestServer"`

	TestRepositories []testRepoConfig.TestRepositoryConfig `json:"TestRepositories"`
	SQLRepositories []sqlRepoConfig.SQLRepositoryConfig `json:"SQLRepositories"`
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
