package hooks

import (
	"fmt"
	"space/backend/models"
	"space/lib/logger"
	"space/lib/server"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// createHandler is handling API calls
func createHandler(ctx *gin.Context) {
	var (
		project models.Project
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

	// validating input
	err = validate.Struct(project)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			sErrors = append(sErrors, err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			wErr := fmt.Errorf("%s, %s", err.Error(), err.Param())
			sErrors = append(sErrors, wErr)
		}
		return
	}

	id, err := (*storage).Create(project.Name)
	if err != nil {
		sErrors = append(sErrors, err)
	}

	res = id
}

func init() {
	server.Authorized.Handle("POST", "/projects", createHandler)
}
