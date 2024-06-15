package authenticator

import (
	// "fmt"
	// "os"

	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	// jwt "github.com/golang-jwt/jwt/v5"
)

// var pubkey []byte = []byte{}

// func validateToken(token string) (bool){

// 	start()
// 	p,err:=jwt.ParseRSAPublicKeyFromPEM(pubkey)
// 	if err!=nil{
// 		fmt.Println("Couldn't parse key")
// 	}
// 	claims := jwt.MapClaims{}

//     tok, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
//         return p, nil
//     })

// 	if err!=nil{
// 		fmt.Println("Couldn't parse token")
// 	}
// 	return tok.Valid

// }

// func start() {
// 	if len(pubkey) ==0 {
// 		var err error
// 		pubkey,err= os.ReadFile("file.cer")
// 		if err!=nil{
// 			fmt.Println("Couldn't read key")
// 		}
// 	}
// }
// IsAuthenticated is a middleware that checks if
// the user has already been authenticated previously.

func checkToken(token string)(int) {
	

	client := &http.Client{}
	req, err := http.NewRequest("GET", provider.UserInfoEndpoint() , nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Bearer " + token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		// fmt.Println("test")
		fmt.Println(err)
		// log.Fatal(err)
	}
	defer resp.Body.Close()
	// bodyText, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return resp.StatusCode
}

func IsAuthenticated(ctx *gin.Context) bool {
	// fmt.Println(sessions.Default(ctx).Get("id_token"))
	// if validateToken(sessions.Default(ctx).Get("id_token").(string)){
	// 	fmt.Println("test")
	// }
	return sessions.Default(ctx).Get("profile") != nil && checkToken(sessions.Default(ctx).Get("access_token").(string)) == http.StatusOK
}


