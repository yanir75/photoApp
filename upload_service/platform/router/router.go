package router

import (
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"update_service/platform/authenticator"
	"update_service/platform/middleware"
	"update_service/platform/uploader"
	"update_service/web/app/callback"
	"update_service/web/app/home"
	"update_service/web/app/login"
	"update_service/web/app/logout"
	"update_service/web/app/user"
	"update_service/web/app/upload"

)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/",middleware.HomePageDecider, home.Handler)
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", middleware.AuthenticatedRedirect, user.Handler)
	router.GET("/logout", logout.Handler)
	router.GET("/upload",middleware.AuthenticatedRedirect, upload.Handler)
	router.POST("/upload",middleware.AuthenticatedRedirect, uploader.Handler)

	// router.POST("/upload",uploader.uploadFile)

	return router
}
