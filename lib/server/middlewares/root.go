package middlewares

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"space/lib/logger"

	"go.uber.org/zap"
)

var (
	environment  = os.Getenv("ENVIRONMENT")
	serverDomain = os.Getenv("SERVER_DOMAIN")
)

func init() {
	environment = strings.ToLower(environment)
}

func SetOrigin(ip net.IP, port int) {
	if environment == "development" {
		Origin = "*"
	} else {
		Origin = fmt.Sprintf("https://%s:%d", serverDomain, port)
	}
}

// middlewareRecovery recovers middleware from a problem
func middlewareRecovery() {
	log := logger.Log.WithOptions(zap.Fields())

	if err := recover(); err != nil {
		_, file, _, _ := runtime.Caller(2)
		file = filepath.Base(file)
		file = strings.Split(file, ".")[0]
		file = strings.Title(file)

		log.Error(fmt.Sprintf("panic recovered in %s Middleware", file),
			zap.String("recover", fmt.Sprintf("%v", err)),
		)
	}
}
