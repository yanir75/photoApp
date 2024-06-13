package upload

import (
	"bytes"
	"encoding/json"
	"net/http"
	"update_service/platform/s3operator"

	// "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Handler for our logged-in user page.
func Handler(ctx *gin.Context) {
	// session := sessions.Default(ctx)
	// profile := session.Get("profile")
	buf := new(bytes.Buffer)

	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	
	err := enc.Encode(s3operator.GenerateUrlMap())
	if err != nil {
		// handle error
	}
	// ctx.String(http.StatusAccepted,"test")
	ctx.HTML(http.StatusOK, "index.html",buf.String())
	// ctx.HTML(http.StatusOK, "index.html", nil)
}
