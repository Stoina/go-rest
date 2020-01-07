package restrepo

import (
	"encoding/json"
	"net/url"
)

// TestRepository exported
// TestRepository ...
type TestRepository struct {
	name string
	url  string
}

// NewTestRepository exported
// NewTestRepository ...
func NewTestRepository(name string, url string) *TestRepository {
	return &TestRepository{
		name: name,
		url:  url}
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
	return repo.name
}

// URL exported
// URL ...
func (repo TestRepository) URL() string {
	return repo.url
}
