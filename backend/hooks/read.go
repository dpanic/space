package hooks

import (
	"errors"
	"space/backend/logic"
	"space/lib/logger"
	"space/lib/server"
	"space/lib/server/middlewares"

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
	defer middlewares.MiddlewareRecovery(ctx)

	var (
		id       = ctx.Param("id")
		readType = "single"
	)
	if id == "" {
		readType = "multiple"
	}

	switch readType {
	case "single":
		// read single result
		logger.Log.Debug("reading project",
			zap.String("id", id),
		)

		obj, err := (*storage).Read(id)
		if err != nil {
			err = errors.New("project doesn't exist")
			sErrors = append(sErrors, err)
			return
		}
		res = obj

	case "multiple":
		// read multiple projects
		logger.Log.Debug("reading multiple projects")

		results, err := (*storage).List()
		if err != nil {
			sErrors = append(sErrors, err)
			return
		}

		logic.List(results)
		res = results
	}
}

func init() {
	server.Authorized.Handle("GET", "/projects/:id", readHandler)

	server.Authorized.Handle("GET", "/projects", readHandler)
}
