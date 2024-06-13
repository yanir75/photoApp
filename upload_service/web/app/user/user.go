package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"update_service/platform/s3operator"

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
	buf := new(bytes.Buffer)

	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	for key,value:= range s3operator.GenerateUrlMap() {
		m["country:"+key] = value
	}
	err := enc.Encode(m)
	if err !=nil {

	}
	// for k, v := range m {
	// 	fmt.Println(k, "=>", v)
	// }
	// ctx.String(200,reflect.ValueOf(profile).String())
	
	ctx.HTML(http.StatusOK, "index.html", buf.String())
}
