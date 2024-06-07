package gallery

import (
	"net/http"
	"encoding/json"
	"update_service/platform/s3operator"
	"github.com/gin-gonic/gin"
	"bytes"
)

// Handler for our home page.
func Handler(ctx *gin.Context) {
	buf := new(bytes.Buffer)

	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	
	err := enc.Encode(s3operator.GenerateUrlMap())
	if err != nil {
		// handle error
	}
	// ctx.String(http.StatusAccepted,"test")
	ctx.HTML(http.StatusOK, "index.html",buf.String())

}
