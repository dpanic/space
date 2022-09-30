package middlewares

import (
	"fmt"
	"path/filepath"
	"runtime"
	"space/lib/logger"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func MiddlewareRecovery(ctx *gin.Context) {
	log := logger.Log.WithOptions(zap.Fields())

	if err := recover(); err != nil {
		_, file, _, _ := runtime.Caller(2)
		file = filepath.Base(file)
		file = strings.Split(file, ".")[0]
		file = strings.Title(file)

		msg := fmt.Sprintf("panic recovered in %s file", file)
		trace := fmt.Sprintf("%v", err)

		ctx.Set("trace", trace)

		log.Error(msg,
			zap.String("trace", trace),
		)
	}
}
