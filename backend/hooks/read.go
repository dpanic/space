package hooks

import (
	"space/lib/logger"
	"space/lib/server"

	"github.com/gin-gonic/gin"
)

// readHandler is handling API calls
func readHandler(ctx *gin.Context) {
	var (
		res     interface{}
		sErrors = make([]error, 0)
	)

	logger.Log.Debug("attempt to read projects")
	defer func() {
		response(ctx, sErrors, res, "read")
	}()

	id := ctx.Param("id")
	if id == "" {
		obj, err := (*storage).Read(id)
		if err != nil {
			sErrors = append(sErrors, err)
			return
		}
		res = obj

	} else {
		results, err := (*storage).List()
		if err != nil {
			sErrors = append(sErrors, err)
			return
		}
		res = results
	}
}

func init() {
	server.Router.Handle("GET", "/projects/:id", readHandler)
	server.Router.Handle("GET", "/projects", readHandler)
}
