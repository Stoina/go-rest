package restserver

// SQLRepositorySettings exported
// ...
type SQLRepositorySettings struct {
	TableName string
	ParameterMapping map[string]string
}

// NewSQLRepositorySettings exported
// ...
func NewSQLRepositorySettings(tableName string, parameterMapping map[string]string) *SQLRepositorySettings {
	return &SQLRepositorySettings{
		TableName: tableName,
		ParameterMapping: parameterMapping};
}