package middlewares

import (
	"os"
	"strings"
)

var (
	environment = os.Getenv("ENVIRONMENT")
)

func init() {
	environment = strings.ToLower(environment)
}
