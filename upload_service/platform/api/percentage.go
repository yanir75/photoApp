package api

import (
	// "net/http"
	// "encoding/json"
	"fmt"
	"net/http"
	"update_service/platform/s3operator"

	"github.com/gin-gonic/gin"
	// "bytes"
	// "fmt"
)

// Handler for our home page.
func PercentageHandler(ctx *gin.Context) {
	param := ctx.Param("filename")
	p,ok := s3operator.PMap.PercentageMap[param]
	if ok {
		perc := p.Perc
		if p.Finished {
			delete(s3operator.PMap.PercentageMap,param)
			fmt.Println(100)
			ctx.String(http.StatusOK,fmt.Sprintf("%v",100));
		} else if perc == 100{
			fmt.Println(99)

			ctx.String(http.StatusOK,fmt.Sprintf("%v",99));
		} else {
			fmt.Println(perc)

			ctx.String(http.StatusOK,fmt.Sprintf("%v",perc));
		}
	} else {
		fmt.Println(0)

		ctx.String(http.StatusOK,fmt.Sprintf("%v",0));
	}


}
