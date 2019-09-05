package utils

import (
	"os"
)

func IsDevEnv() bool {
	return os.Getenv("APP_ENV") == "development"
}
