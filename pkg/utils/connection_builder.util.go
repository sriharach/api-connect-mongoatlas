package utils

import (
	"fmt"
	"os"
)

type ConnectTypeEnum string

const (
	POSTGRES ConnectTypeEnum = "POSTGRES"
	MYSQL    ConnectTypeEnum = "MYSQL"
	REDIS    ConnectTypeEnum = "REDIS"
	FIBER    ConnectTypeEnum = "FIBER"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n ConnectTypeEnum) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case POSTGRES:
		// URL for PostgreSQL connection.
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SSL_MODE"),
		)
	case MYSQL:
		// URL for Mysql connection.
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
	case REDIS:
		// URL for Redis connection.
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		)
	case FIBER:
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}
