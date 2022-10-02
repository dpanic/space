package hooks

import (
	"space/backend/logic"
	"space/backend/models"
	"space/lib/logger"
	"space/lib/server"
	"space/lib/server/middlewares"

	"github.com/gin-gonic/gin"
)

// createHandler is handling API calls
func createHandler(ctx *gin.Context) {
	var (
		project *models.Project
		res     interface{}
		isRaw   bool
		sErrors = make([]error, 0)
	)

	logger.Log.Debug("attempt to create project")
	defer func() {
		response(ctx, sErrors, res, isRaw, "create")
	}()
	defer middlewares.MiddlewareRecovery(ctx)

	// initial error, which is removed later
	err := ctx.ShouldBind(&project)
	if err != nil {
		sErrors = append(sErrors, err)
		return
	}

	// execute logic
	newProject, err := logic.Create(project.Name)
	if err != nil {
		sErrors = append(sErrors, err)
		return
	}

	// update storage
	err = (*storage).Create(newProject)
	if err != nil {
		sErrors = append(sErrors, err)
		return
	}

	res = newProject.Id
}

func init() {
	server.Authorized.Handle("POST", "/projects", createHandler)
}
