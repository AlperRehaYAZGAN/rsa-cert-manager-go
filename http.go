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

	// GET /resources - get all resources
	r.GET("/resources", controllers.GetResourcesHandler)
	// POST /resources - create resource
	r.POST("/resources", controllers.CreateResourceHandler)
	// PUT /resources/:id - update resource
	r.PUT("/resources/:id", controllers.UpdateResourceHandler)
	// DELETE /resources/:id - delete resource
	r.DELETE("/resources/:id", controllers.DeleteResourceHandler)

	// GET /platforms - get all platforms
	r.GET("/platforms", controllers.GetPlatformsHandler)
	// POST /platforms - create platform
	r.POST("/platforms", controllers.CreatePlatformHandler)
	// PUT /platforms/:id - update platform
	r.PUT("/platforms/:id", controllers.UpdatePlatformHandler)
	// DELETE /platforms/:id - delete platform
	r.DELETE("/platforms/:id", controllers.DeletePlatformHandler)

	// GET /services - get all services
	r.GET("/services", controllers.GetServicesHandler)
	// POST /services - create service
	r.POST("/services", controllers.CreateServiceHandler)
	// PUT /services/:id - update service
	r.PUT("/services/:id", controllers.UpdateServiceHandler)
	// DELETE /services/:id - delete service
	r.DELETE("/services/:id", controllers.DeleteServiceHandler)

	// GET /scans - get all scans
	r.GET("/scans", controllers.GetScansHandler)
	// DELETE /scans/:id - delete scan
	r.DELETE("/scans/:id", controllers.DeleteScanHandler)

}

// HelloWorldIndexHandler is a simple health check endpoint
func HelloWorldIndexHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "World",
	})
}
