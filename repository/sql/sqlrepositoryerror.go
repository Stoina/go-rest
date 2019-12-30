package rest

// RepositoryError exported
// RepositoryError ...
type RepositoryError struct {
	ErrorOccurred bool
	ErrorMsg      string
}

// NewError exported
// NewError ...
func NewError(err error) *RepositoryError {
	return &RepositoryError{
		ErrorOccurred: true,
		ErrorMsg:      err.Error()}
}
