package restrepo

import "encoding/json"

// RepositoryResult exported
// RepositoryResult ...
type RepositoryResult struct {
	Successful      bool
	Error           *RepositoryError
	Data            string
	ResponseMessage string
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
