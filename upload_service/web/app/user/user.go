package user

import (
	"encoding/json"
	"net/http"
	// "reflect"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Handler for our logged-in user page.
func Handler(ctx *gin.Context) {
	
	session := sessions.Default(ctx)
	profile := session.Get("profile")
	m, ok := profile.(map[string]interface{})
	if ok {
		m["permissions"] = m[os.Getenv("URL")]
	}
	// for k, v := range m {
	// 	fmt.Println(k, "=>", v)
	// }
	// ctx.String(200,reflect.ValueOf(profile).String())
	j,_ := json.Marshal(profile)
	ctx.HTML(http.StatusOK, "index.html", string(j))
}
