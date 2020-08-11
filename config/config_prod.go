// +build prod

package config

const (
	DB_USER     = "admin"
	DB_PASSWORD = "admin"
	DB_DATABASE = "forum_db"
	DB_HOST     = "127.0.0.1"
	//DB_HOST=127.0.0.1                # RUNNING THE APP WITHOUT DOCKER
	DB_PORT    = 3306
	API_PORT   = 8080
	API_SECRET = "98hbun98h"
)
