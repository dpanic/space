package hooks

import (
	"space/lib/logger"
	"space/lib/server"
	"space/lib/server/middlewares"

	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

// deleteHandler is handling API calls
func deleteHandler(ctx *gin.Context) {
	var (
		sErrors = make([]error, 0)
		res     interface{}
		isRaw   bool
		id      string
	)

	logger.Log.Debug("attempt to delete project")
	defer func() {
		response(ctx, sErrors, res, isRaw, "delete")
	}()
	defer middlewares.MiddlewareRecovery(ctx)

	id = ctx.Param("id")
	if id == "" {
		err := errors.New("id is empty")
		sErrors = append(sErrors, err)
		return
	}

	_, err := (*storage).Read(id)
	if err != nil {
		err = errors.New("project doesn't exist")
		sErrors = append(sErrors, err)
		return
	}

	err = (*storage).Delete(id)
	if err != nil {
		wErr := fmt.Errorf("error in deleting project: %s", err.Error())
		sErrors = append(sErrors, wErr)
		return
	}

	res = id
}

func init() {
	server.Authorized.Handle("DELETE", "/projects/:id", deleteHandler)
}
