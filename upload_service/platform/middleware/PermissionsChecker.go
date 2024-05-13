package middleware

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func contains(li []interface{}, perm string) bool {
	for _, item := range li {
		if item == perm {
			return true
		}
	}
	return false
}

func containsList(li1 []interface{}, li2 []string) bool {
	for _, item := range li2 {
		if !contains(li1, item) {
			return false
		}
	}
	return true
}

func checkPermissions(profile interface{}, permissions []string) (map[string]interface{}, bool) {
	info, ok := profile.(map[string]interface{})
	authorized := false
	if ok {
		if permissionsList, err := info[os.Getenv("URL")].([]interface{}); err {
			info["permissions"] = permissionsList
			if containsList(permissionsList, permissions) {
				authorized = true
			}
		}
	}
	return info, authorized
}

// IsAuthenticated is a middleware that checks if
// the user has already been authenticated previously.
func permissionsChecker(ctx *gin.Context,permissions []string) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")
	// m, ok := profile.(map[string]interface{})
	// authorized := false
	// if ok  {
	// 	if permissionsList,err := m[os.Getenv("URL")].([]interface{}); err{
	// 		m["permissions"] = permissionsList
	// 		if slices.Contains(permissionsList,"upload"){
	// 			authorized = true
	// 			ctx.Next()
	// 		}
	// 	}
	// }
	// info,authorized := checkPermissions(profile,["upload"])
	if info, authorized := checkPermissions(profile, permissions); authorized {
		ctx.Next()
	} else {
		ctx.HTML(http.StatusUnauthorized, "user.html", info)
		ctx.Abort()
	}

}

func PermissionsHandler(permissions []string) gin.HandlerFunc {
    fn := func(c *gin.Context) {
        // Your handler code goes in here - e.g.
        permissionsChecker(c,permissions)
    }

    return gin.HandlerFunc(fn)
}