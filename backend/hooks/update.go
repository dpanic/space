package hooks

import (
	"errors"
	"space/backend/logic"
	"space/backend/models"
	"space/lib/logger"
	"space/lib/server"

	"github.com/gin-gonic/gin"
)

// updateHandler is handling API calls
func updateHandler(ctx *gin.Context) {
	var (
		sErrors    = make([]error, 0)
		res        interface{}
		newProject *models.Project
	)

	logger.Log.Debug("attempt to update project")
	defer func() {
		response(ctx, sErrors, res, "update")
	}()

	id := ctx.Param("id")
	if id == "" {
		err := errors.New("id is empty")
		sErrors = append(sErrors, err)
		return
	}

	err := ctx.ShouldBind(&newProject)
	if err != nil {
		sErrors = append(sErrors, err)
		return
	}

	currentProject, err := (*storage).Read(id)
	if err != nil {
		err = errors.New("project doesn't exist")
		sErrors = append(sErrors, err)
		return
	}

	// execute logic
	updatedProject, err := logic.Update(currentProject, newProject)
	if err != nil {
		sErrors = append(sErrors, err)
		return
	}

	// update storage
	_, err = (*storage).Update(id, updatedProject)
	if err != nil {
		sErrors = append(sErrors, err)
		return
	}

	res = updatedProject
}

func init() {
	server.Authorized.Handle("PUT", "/projects/:id", updateHandler)
}
