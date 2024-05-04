package authenticator

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// IsAuthenticated is a middleware that checks if
// the user has already been authenticated previously.
func IsAuthenticated(ctx *gin.Context) bool {
	return sessions.Default(ctx).Get("profile") != nil
}
