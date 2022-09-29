package hooks

import (
	"fmt"
	"os"
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
	rootDir        string
	desiredStorage = "disk"
	storage        *storages.Storage
)

func Bind() {
}

func response(ctx *gin.Context, sErrors []error, res interface{}, action string) {
	var (
		out = make(map[string]interface{})
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

		out["status"] = false
		out["message"] = msg
		out["errors"] = parsedErrors

	} else {
		msg := fmt.Sprintf("success in %s", action)
		logger.Log.Info(msg)

		out = make(map[string]interface{})
		out["status"] = true
		out["message"] = msg
		out["data"] = res
	}

	res = out
	serverAction.Response(ctx, res)
}

func init() {
	rootDir = os.Getenv("rootDir")
	if rootDir == "" {
		rootDir = "/tmp"
	}

	var found bool
	storage, found = storages.Get(desiredStorage)
	if !found {
		logger.Log.Panic("desired storage not found",
			zap.String("storage", desiredStorage))
		return
	}
}
