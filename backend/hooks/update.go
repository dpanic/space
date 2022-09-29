package hooks

import (
	"errors"
	"space/backend/models"
	"space/lib/logger"
	"space/lib/server"

	"github.com/gin-gonic/gin"
)

// updateHandler is handling API calls
func updateHandler(ctx *gin.Context) {
	var (
		sErrors = make([]error, 0)
		project models.Project
	)

	logger.Log.Debug("attempt to update project")
	defer func() {
		response(ctx, sErrors, nil, "update")
	}()

	id := ctx.Param("id")
	if id == "" {
		err := errors.New("id is empty")
		sErrors = append(sErrors, err)
		return
	}

	err := ctx.ShouldBind(&project)
	if err != nil {
		sErrors = append(sErrors, err)
		return
	}

	_, err = (*storage).Read(id)
	if err != nil {
		err = errors.New("project doesn't exist")
		sErrors = append(sErrors, err)
		return
	}

	err = (*storage).Update(id, &project.Data)
	if err != nil {
		sErrors = append(sErrors, err)
		return
	}
}

func init() {
	server.Authorized.Handle("PUT", "/projects/:id", updateHandler)
}
