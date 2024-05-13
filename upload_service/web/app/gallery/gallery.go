package gallery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"update_service/platform/s3operator"
)


// Handler for our home page.
func Handler(ctx *gin.Context) {
	data := map[string]interface{}{
		"Folders": s3operator.GetS3Folders(),
	}
	// ctx.String(http.StatusAccepted,"test")
	ctx.HTML(http.StatusOK, "index.html",data)

}
