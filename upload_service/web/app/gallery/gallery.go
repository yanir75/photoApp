package gallery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"update_service/platform/uploader"
)

// Handler for our home page.
func Handler(ctx *gin.Context) {
	uploader.GetS3Folders()
	ctx.String(http.StatusAccepted,"test")
}
