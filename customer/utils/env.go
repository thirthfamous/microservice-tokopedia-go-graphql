package utils

import "os"

func InitEnvironment() {
	os.Setenv("DB_URL", "127.0.0.1:3306")
	os.Setenv("DB_NAME", "user")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "")
}
