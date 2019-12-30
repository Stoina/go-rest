package rest

import (
	"net/url"

	db "github.com/Stoina/go-database"
)

// SQLRepository exported
// SQLRepository ...
type SQLRepository struct {
	name         string
	url          string
	dbConnection *db.Connection
	config       *SQLRepositoryConfig
}

// NewSQLRepository exported
// NewSQLRepository ...
func NewSQLRepository(name string, url string, dbConn *db.Connection, config *SQLRepositoryConfig) *SQLRepository {
	return &SQLRepository{
		name:         name,
		url:          url,
		dbConnection: dbConn,
		config:       config}
}

// Add exported
// Add ...
func (sqlRepo SQLRepository) Add(jsonInput string) {

}

// Get exported
// Get ...
func (sqlRepo SQLRepository) Get(calledURL *url.URL) string {

	repoResult := RepositoryResult{}

	queryResult, err := sqlRepo.dbConnection.Query(sqlRepo.config.Statement)

	if err != nil {
		repoResult.Error = NewError(err)
		return repoResult.ConvertToJSON()
	}

	data, err := queryResult.ConvertToJSON()

	if err != nil {
		repoResult.Error = NewError(err)
		return repoResult.ConvertToJSON()
	}

	repoResult.Data = data
	return repoResult.ConvertToJSON()
}

// Name exported
// Name ....
func (sqlRepo SQLRepository) Name() string {
	return sqlRepo.name
}

// URL exported
// URL ...
func (sqlRepo SQLRepository) URL() string {
	return sqlRepo.url
}
