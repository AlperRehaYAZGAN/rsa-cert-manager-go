/**
*	This is the Platform CRUD Controller.
 */
package controllers

import (
	// built-in
	"net/http"

	// web framework
	"github.com/gin-gonic/gin"
)

func GetPlatformsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"type":   "ttrn:::api/platforms/get",
		"platforms": []string{
			"/api/services/generate/",
		},
	})
	return
}

func CreatePlatformHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"status": true,
		"type":   "ttrn:::api/platforms/create",
		"platform": []string{
			"/api/services/platforms",
		},
	})
	return
}

func UpdatePlatformHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"type":   "ttrn:::api/platforms/update",
		"platform": []string{
			"/api/services/platforms",
		},
	})
	return
}

func DeletePlatformHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"type":   "ttrn:::api/services/delete",
		"platform": []string{
			"/api/platforms/paas",
		},
	})
	return
}
