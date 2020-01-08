package restserver

import (
	"net/url"
)

// Repository exported
// Repository ...
type Repository interface {
	Post(par string) *RepositoryResult
	Put(par string) *RepositoryResult
	Patch(par string) *RepositoryResult
	Delete(par string) *RepositoryResult
	Get(calledURL *url.URL) *RepositoryResult
	Name() string 
	URL() string
}