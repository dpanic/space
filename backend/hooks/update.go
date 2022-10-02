package hooks

import (
	"errors"
	"space/backend/logic"
	"space/backend/models"
	"space/lib/logger"
	"space/lib/server"
	"space/lib/server/middlewares"
	"strings"

	"github.com/gin-gonic/gin"
)

// updateHandler is handling API calls
func updateHandler(ctx *gin.Context) {
	var (
		sErrors    = make([]error, 0)
		res        interface{}
		isRaw      bool
		newProject *models.Project
	)

	logger.Log.Debug("attempt to update project")
	defer func() {
		response(ctx, sErrors, res, isRaw, "update")
	}()
	defer middlewares.MiddlewareRecovery(ctx)

	err := ctx.ShouldBind(&newProject)
	if err != nil {
		sErrors = append(sErrors, err)
		return
	}

	id := ctx.Param("id")
	if id == "" {
		err := errors.New("id is empty")
		sErrors = append(sErrors, err)
		return
	}

	outType, _ := ctx.GetQuery("o")
	outType = strings.ToLower(outType)

	currentProject, err := (*storage).Read(id)
	if err != nil {
		err = errors.New("project doesn't exist")
		sErrors = append(sErrors, err)
		return
	}

	// execute logic
	updatedProject, updateErrors := logic.Update(currentProject, newProject)
	if updateErrors != nil {
		sErrors = append(sErrors, updateErrors...)
		return
	}

	// update storage
	_, err = (*storage).Update(id, updatedProject)
	if err != nil {
		sErrors = append(sErrors, err)
		return
	}

	if outType == "geojson" {
		res = updatedProject.Data.BuildingSplits
		isRaw = true
	} else {
		res = updatedProject
	}
}

func init() {
	server.Authorized.Handle("PUT", "/projects/:id", updateHandler)
}
