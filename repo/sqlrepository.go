package restserver

import (
	"net/url"

	db "github.com/Stoina/go-database"
)

// SQLRepository exported
// SQLRepository ...
type SQLRepository struct {
	name         	string
	url          	string
	dbConnection 	*db.Connection
}

// NewSQLRepository exported
// NewSQLRepository ...
func NewSQLRepository(name string, url string, dbConn *db.Connection) *SQLRepository {
	return &SQLRepository{
		name:         name,
		url:          url,
		dbConnection: dbConn}
}

// Post exported
// With a post request to the container resource you can create a new resource.
func (sqlRepo SQLRepository) Post(par string) *RepositoryResult {
	return nil
}

// Put exported
// With a put request to the container resource you can overwrite the resource with the representation in the request.
func (sqlRepo SQLRepository) Put(par string) *RepositoryResult {
	return nil
}

// Patch exported
// With the http patch method, individual properties of a resource can be manipulated in a targeted manner.
func (sqlRepo SQLRepository) Patch(par string) *RepositoryResult {
	return nil
}

// Delete exported
// With this method an existing resource can be deleted
func (sqlRepo SQLRepository) Delete(par string) *RepositoryResult {
	return nil
}

// Get exported
// Get ...
func (sqlRepo SQLRepository) Get(calledURL *url.URL) *RepositoryResult {
	return sqlRepo.Query("")
}

// Query exported
// Query ...
func (sqlRepo *SQLRepository) Query(query string) *RepositoryResult {
	repoResult := RepositoryResult{}

	queryResult, err := sqlRepo.dbConnection.Query(query)

	if err != nil {
		repoResult.ErrorOccurred = true
		repoResult.ErrorMessage = err.Error()

		return &repoResult
	}

	data, err := queryResult.ConvertToJSON()

	if err != nil {
		repoResult.ErrorOccurred = true
		repoResult.ErrorMessage = err.Error()
		
		return &repoResult
	}

	repoResult.Data = data

	return &repoResult
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
