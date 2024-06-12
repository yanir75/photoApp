package api

import (
	// "net/http"
	// "encoding/json"
	// "update_service/platform/s3operator"
	"bytes"
	"encoding/json"
	"net/http"
	"update_service/platform/s3operator"

	"github.com/gin-gonic/gin"
	// "bytes"
	// "fmt"
)

// Handler for our home page.
func CountryHandler(ctx *gin.Context) {
	buf := new(bytes.Buffer)

	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	
	err := enc.Encode(s3operator.GenerateUrlCountryMap(ctx.Param("country")))
	if err != nil {
		// handle error
	}
	// ctx.String(http.StatusAccepted,"test")
	ctx.JSON(http.StatusOK,buf.String())

}
