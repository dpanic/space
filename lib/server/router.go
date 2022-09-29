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

	// Define and set request ID
	Router.Use(func(ctx *gin.Context) {
		id, _ := crypto.UUID()
		ctx.Set("id", id)
		ctx.Next()
	})

	Router.Use(middlewares.Stats())
	Router.Use(middlewares.NoCache())
	Router.Use(middlewares.Session())
	Router.Use(middlewares.CORS())
	Router.Use(middlewares.Security())
}
