package restrepoconfig

// SQLRepositoryConfig exported
// SQLRepositoryConfig ...
type SQLRepositoryConfig struct {
	Name string `json:"Name"`
	URL string `json:"URL"`
	DatabaseConfig DatabaseConfig `json:"DBSettings"`
	Statement string
}

// DatabaseConfig exported
// DatabaseConfig ...
type DatabaseConfig struct {
	Host     string `json:"Server"`
	Port	 string `json:"Port"`
	DBName   string `json:"Database"`
	Username string `json:"Username"`
	Password string `json:"Password"`
} 

