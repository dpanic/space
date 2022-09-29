package hooks

import (
	"fmt"
	"space/lib/server"
	serverAction "space/lib/server/action"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	started time.Time
)

//go:generate easytags $GOFILE json:camel
type HealthResponse struct {
	Message string `json:"message"`
	Uptime  string `json:"uptime"`
}

func healthHandler(ctx *gin.Context) {
	ctx.Status(200)

	res := HealthResponse{
		Message: "all ok!",
		Uptime:  fmt.Sprintf("%v", time.Since(started)),
	}
	serverAction.Response(ctx, res)
}

func init() {
	started = time.Now()

	server.Router.Handle("ANY", "/", healthHandler)
	server.Router.NoRoute(healthHandler)
}
