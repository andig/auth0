package router

import (
	"encoding/gob"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"01-Login/platform/authenticator"
	"01-Login/platform/middleware"
	"01-Login/web/app/callback"
	"01-Login/web/app/login"
	"01-Login/web/app/logout"
	"01-Login/web/app/user"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))

	store.Options(sessions.Options{
		Path:     "/",   // Available across the entire domain
		Domain:   "",    // Adjust to your domain, use an empty string for localhost or omit for default behavior
		MaxAge:   3600,  // Expires in 1 hour
		Secure:   false, // Set to true if serving over HTTPS, false otherwise
		HttpOnly: true,  // Recommended: JavaScript can't access the cookie
	})

	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", middleware.IsAuthenticated, user.Handler)
	router.GET("/logout", logout.Handler)

	return router
}
