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
		"intent": "ttrn:::api/platforms/get",
		"type":   "get-platforms",
		"platforms": []string{
			"/api/services/generate/",
		},
	})
	return
}

func CreatePlatformHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"status": true,
		"intent": "ttrn:::api/platforms/create",
		"type":   "create-platform",
		"platform": []string{
			"/api/services/platforms",
		},
	})
	return
}

func UpdatePlatformHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"intent": "ttrn:::api/platforms/update",
		"type":   "update-platform",
		"platform": []string{
			"/api/services/platforms",
		},
	})
	return
}

func DeletePlatformHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"intent": "ttrn:::api/platforms/delete",
		"type":   "delete-platform",
		"platform": []string{
			"/api/platforms/paas",
		},
	})
	return
}
