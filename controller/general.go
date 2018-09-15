package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RouteSetup will configure the basic routes for this system
func RouteSetup(r *gin.Engine) {
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

}

// HandlerRoot godoc
// @Summary Show a Health check response
// @Description get health check of system
// @Accept  json
// @Produce  json
// @Success 200 string string
// @Router / [get]
func HandlerRoot(c *gin.Context) {
	//Do DB health check
	//Do redis health check
	//Do any other external dependency health check
	//If all ok, respond "OK", else respond "NOT_OK"
	c.JSON(http.StatusOK, gin.H{"healthcheck": "OK"})
}
