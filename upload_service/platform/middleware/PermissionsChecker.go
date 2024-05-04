package middleware

import (
	"slices"
	"net/http"
	"os"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// IsAuthenticated is a middleware that checks if
// the user has already been authenticated previously.
func PermissionsChecker(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")
	m, ok := profile.(map[string]interface{})
	authorized := false
	if ok  {
		if permissionsList,err := m[os.Getenv("URL")].([]interface{}); err{
			m["permissions"] = permissionsList
			if slices.Contains(permissionsList,"upload"){
				authorized = true
				ctx.Next()
			}
		} 
	}	


	if !authorized {
		ctx.HTML(http.StatusUnauthorized, "user.html", m)
	}

}