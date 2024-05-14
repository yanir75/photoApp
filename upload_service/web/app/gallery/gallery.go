package gallery

import (
	"net/http"

	"encoding/json"
	"update_service/platform/s3operator"

	"github.com/gin-gonic/gin"
)

// Handler for our home page.
func Handler(ctx *gin.Context) {

	strBytes,err := json.Marshal(s3operator.GenerateUrlMap())
	if err != nil {
		// handle error
	}
	// ctx.String(http.StatusAccepted,"test")
	ctx.HTML(http.StatusOK, "index.html",string(strBytes))

}
