package config

var jwtSecretKey = "udia#y387dyJkanadk7&54"

type DatabaseConfig struct {
	Host               string
	Port               string
	User               string
	Password           string
	Name               string
	Charset            string
	MaxConnections     int
	MaxIdleConnections int
}

func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:               "127.0.0.1",
		Port:               "3306",
		User:               "root",
		Password:           "admin123",
		Name:               "myblog",
		Charset:            "utf8mb4",
		MaxConnections:     10,
		MaxIdleConnections: 5,
	}
}

func GetJWTSecretKey() string {
	return jwtSecretKey
}
