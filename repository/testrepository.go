package rest

import (
	"encoding/json"
	"net/url"
)

// TestRepository exported
// TestRepository ...
type TestRepository struct {
	URLString  string
	NameString string
}

// Add exported
// Add ...
func (repo TestRepository) Add(jsonInput string) {

}

// Get exported
// Get ...
func (repo TestRepository) Get(calledURL *url.URL) string {
	jsonBytes, _ := json.Marshal("Test")
	return string(jsonBytes)
}

// Name exported
// Name ....
func (repo TestRepository) Name() string {
	return repo.NameString
}

// URL exported
// URL ...
func (repo TestRepository) URL() string {
	return repo.URLString
}
