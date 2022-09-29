package middlewares

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"space/lib/logger"

	"go.uber.org/zap"
)

var (
	environment = os.Getenv("ENVIRONMENT")
)

func init() {
	environment = strings.ToLower(environment)
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
