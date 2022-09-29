package hooks

import (
	"space/lib/logger"
	"space/lib/server"

	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

// deleteHandler is handling API calls
func deleteHandler(ctx *gin.Context) {
	var (
		sErrors = make([]error, 0)
		id      string
	)

	logger.Log.Debug("attempt to delete project")
	defer func() {
		response(ctx, sErrors, nil, "delete")
	}()

	id = ctx.Param("id")
	if id == "" {
		err := errors.New("id is empty")
		sErrors = append(sErrors, err)
		return
	}

	err := (*storage).Delete(id)
	if err != nil {
		wErr := fmt.Errorf("error in deleting project: %s", err.Error())
		sErrors = append(sErrors, wErr)
		return
	}
}

func init() {
	server.Router.Handle("DELETE", "/project/:id", deleteHandler)
}