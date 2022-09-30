package server

import (
	"space/lib/crypto"
	"space/lib/server/middlewares"

	"github.com/gin-gonic/gin"
)

var (
	Router     = gin.New()
	Authorized *gin.RouterGroup
	user       = "space"
	pass       = "maker"
)

func init() {
	Authorized = Router.Group("/", gin.BasicAuth(gin.Accounts{
		user: pass,
	}))

	Authorized.Use(func(ctx *gin.Context) {
		id, _ := crypto.UUID()
		ctx.Set("id", id)
		ctx.Next()
	})

	Authorized.Use(middlewares.Stats())
	Authorized.Use(middlewares.NoCache())
	Authorized.Use(middlewares.Session())
	Authorized.Use(middlewares.CORS())
	Authorized.Use(middlewares.Security())
}
