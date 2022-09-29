package middlewares

import (
	"space/lib/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Stats() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		id := ctx.GetString("id")

		log := logger.Log.WithOptions(zap.Fields(
			zap.String("id", id),
			zap.String("ip", ip),
			zap.String("method", ctx.Request.Method),
			zap.String("uri", ctx.Request.RequestURI),
		),
		)

		log.Debug("request started")

		tsStart := time.Now()
		defer middlewareRecovery()
		ctx.Next()

		log.Debug("request finished",
			zap.Duration("duration", time.Since(tsStart)),
		)
	}
}
