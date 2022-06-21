package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(param string) string {
	err := godotenv.Load()
	PanicIfError(err, "Error loading .env file")

	return os.Getenv(param)
}

func PanicIfError(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}
