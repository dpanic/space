package hooks

import (
	"fmt"
	"space/lib/logger"
	"space/lib/server"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	if id != "" {
		logger.Log.Debug("reading project",
			zap.String("id", id),
		)

		obj, err := (*storage).Read(id)
		if err != nil {
			sErrors = append(sErrors, err)
			return
		}
		res = obj

	} else {
		logger.Log.Debug("reading multiple projects")

		results, err := (*storage).List()
		if err != nil {
			sErrors = append(sErrors, err)
			return
		}

		fmt.Println(1)
		res = results
	}
}

func init() {
	server.Authorized.Handle("GET", "/projects/:id", readHandler)
	server.Authorized.Handle("GET", "/projects", readHandler)
}
