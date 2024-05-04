package middleware

import (
	"net/http"

	"update_service/platform/authenticator"

	"github.com/gin-gonic/gin"
)

// IsAuthenticated is a middleware that checks if
// the user has already been authenticated previously.
func AuthenticatedRedirect(ctx *gin.Context) {
	if !authenticator.IsAuthenticated(ctx) {
		ctx.Redirect(http.StatusSeeOther, "/")
	} else {
		ctx.Next()
	}
}
