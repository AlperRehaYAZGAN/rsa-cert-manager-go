package main

import (
	// built-in
	"net/http"

	// local packages
	controllers "ttcertman/src/controllers"

	// web framework
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/", HelloWorldIndexHandler)

	// GET /generate/ssl - generate ssl key pair
	r.GET("/generate/ssl", controllers.GenerateSSLHandler)
	// GET /generate/ssh - generate ssh key pair
	r.GET("/generate/ssh", controllers.GenerateSSHHandler)

	// GET /scan/ssl - scan ssl certificate
	r.GET("/scan/ssl", controllers.ScanSSLHandler)
}

// HelloWorldIndexHandler is a simple health check endpoint
func HelloWorldIndexHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "World",
	})
}
