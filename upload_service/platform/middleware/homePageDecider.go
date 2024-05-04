package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"update_service/platform/authenticator"
)

// IsAuthenticated is a middleware that checks if
// the user has already been authenticated previously.
func HomePageDecider(ctx *gin.Context) {
	if authenticator.IsAuthenticated(ctx) {
		ctx.Redirect(http.StatusSeeOther, "/upload")
	} else {
		ctx.Next()
	}
}
