package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter configura tutte le rotte
func SetupRouter(r *gin.Engine) {
	// // Servire i file statici (CSS, JS)
	// r.Static("/static", "./static")

	// // Servire i file HTML
	// r.LoadHTMLGlob("views/*")

	// // Rotte per le pagine web
	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "login.html", nil)
	// })

	// r.GET("/register", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "register.html", nil)
	// })

	// r.GET("/home", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "home.html", nil)
	// })

	// Includi le rotte separate
	SetupUserRoutes(r)
	// SetupAlbumRoutes(r)
	// SetupCartRoutes(r)
	SetupChatCompletionRoutes(r)
	SetupMessageRoutes(r)
}
