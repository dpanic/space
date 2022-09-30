package hooks

import (
	"fmt"
	"space/backend/models"
	"space/backend/storages"
	"space/lib/logger"
	serverAction "space/lib/server/action"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// use a single instance of Validate, it caches struct info
var (
	validate       = validator.New()
	storage        *storages.Storage
	desiredStorage = "disk"
)

func Bind() {
}

//go:generate easytags $GOFILE json:snake
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Total   int         `json:"total"`
	Results interface{} `json:"results"`
	Errors  []string    `json:"errors,omitempty"`
}

func response(ctx *gin.Context, sErrors []error, res interface{}, action string) {
	var (
		response Response
	)

	if len(sErrors) > 0 {
		var (
			errMessage   = ""
			parsedErrors = make([]string, 0)
		)
		for _, err := range sErrors {
			errMessage += err.Error() + " "
			parsedErrors = append(parsedErrors, err.Error())
		}

		err := fmt.Errorf(errMessage)

		msg := fmt.Sprintf("failed to %s", action)
		logger.Log.Error(msg,
			zap.Error(err),
		)

		response.Status = false
		response.Message = msg
		response.Errors = parsedErrors

	} else {
		msg := fmt.Sprintf("success in %s", action)
		value := fmt.Sprintf("%+v", res)
		logger.Log.Info(msg,
			zap.String("res", value),
		)

		response.Status = true
		response.Message = msg
		response.Results = res

		total := 0
		switch res := res.(type) {
		case []interface{}:
			total = len(res)

		case []string:
			total = len(res)

		case []*models.Project:
			total = len(res)
		}

		response.Total = total

	}

	serverAction.Response(ctx, response)
}

func init() {
	var found bool
	storage, found = storages.Get(desiredStorage)
	if !found {
		logger.Log.Panic("desired storage not found",
			zap.String("storage", desiredStorage))
		return
	}
}
