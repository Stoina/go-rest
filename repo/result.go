package restserver

import "encoding/json"

// RepositoryResult exported
// RepositoryResult ...
type RepositoryResult struct {
	Data            string
	ErrorOccurred 	bool
	ErrorMessage 	string
	ResponseMessage string
	Successful      bool
}

// NewRepositoryResult exported
// NewRepositoryResult ...
func NewRepositoryResult(data string, errorOccurred bool, errorMessage string, responseMessage string, successful bool) *RepositoryResult {
	return &RepositoryResult {
		Data: data,
		ErrorOccurred: errorOccurred,
		ErrorMessage: errorMessage,
		ResponseMessage: responseMessage,
		Successful: successful}
}

// ConvertToJSON exported
// ConvertToJSON ...
func (repoResult *RepositoryResult) ConvertToJSON() string {
	converted, err := json.Marshal(repoResult)

	if err != nil {
		return ""
	}

	return string(converted)
}
