package gallery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"update_service/platform/s3operator"
)

// Handler for our home page.
func Handler(ctx *gin.Context) {
	s3operator.GetS3Folders()
	ctx.String(http.StatusAccepted,"test")
}
