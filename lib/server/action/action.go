package action

import (
	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// encodeResponse in proper formatted output
func encodeResponse(res interface{}) (out string, err error) {
	b := new(bytes.Buffer)
	enc := json.NewEncoder(b)

	// NOTES: escape XSS
	// we can do this manually by escaping: < > ' " &
	enc.SetEscapeHTML(true)

	enc.SetIndent("", "    ")
	err = enc.Encode(res)
	out = b.String()

	return
}

// Response performs response on stdout
func Response(ctx *gin.Context, res interface{}) {
	ctx.Header("Content-Type", "application/json; charset=utf-8")

	obj, _ := encodeResponse(res)
	statusCode := ctx.Writer.Status()
	ctx.Data(statusCode, "application/json; charset=utf-8", []byte(obj))
}
