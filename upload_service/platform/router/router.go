package router

import (
	"encoding/gob"
	// "fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"update_service/platform/authenticator"
	"update_service/platform/middleware"
	"update_service/platform/s3operator"
	"update_service/web/app/callback"
	"update_service/web/app/gallery"
	"update_service/web/app/home"
	"update_service/web/app/login"
	"update_service/web/app/logout"
	"update_service/web/app/upload"
	"update_service/web/app/user"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))
	router.Static("/gallery", "client/dist")

	router.Static("/public", "web/static")

	
	// router.Handle("/test","hello-world/build")
		// router.Use(static.Serve("/test",static.LocalFile("hello-world/build",true)))
	// router.LoadHTMLGlob("web/template/*")
	router.LoadHTMLFiles("client/dist/index.html","web/template/user.html","web/template/upload.html")


	router.GET("/", middleware.HomePageDecider, home.Handler)
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", middleware.AuthenticatedRedirect, user.Handler)
	router.GET("/logout", logout.Handler)
	router.GET("/gallery", gallery.Handler)
	router.GET("/upload", middleware.AuthenticatedRedirect, middleware.PermissionsHandler([]string {"upload"}), upload.Handler)
	router.POST("/upload", middleware.AuthenticatedRedirect, middleware.PermissionsHandler([]string {"upload"}), s3operator.Handler)

	// router.POST("/upload",uploader.uploadFile)

	return router
}
