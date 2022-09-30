package hooks

import (
	"space/backend/logic"
	"space/backend/models"
	"space/lib/logger"
	"space/lib/server"

	"github.com/gin-gonic/gin"
)

// createHandler is handling API calls
func createHandler(ctx *gin.Context) {
	var (
		project *models.Project
		res     interface{}
		sErrors = make([]error, 0)
	)

	logger.Log.Debug("attempt to create project")
	defer func() {
		response(ctx, sErrors, res, "create")
	}()

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
