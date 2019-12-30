package rest

import (
	"net/url"
)

// Repository exported
// Repository ...
type Repository interface {
	Add(jsonInput string)
	Get(calledURL *url.URL) string
	Name() string
	URL() string
}
