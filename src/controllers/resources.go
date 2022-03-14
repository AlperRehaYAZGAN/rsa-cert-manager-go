/**
*	This is the Resource CRUD Controller.
 */
package controllers

import (
	// built-in
	"net/http"

	// web framework
	"github.com/gin-gonic/gin"
)

func GetResourcesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"type":   "ttrn:::api/resources/get",
		"resources": []string{
			"/api/services/generate/",
		},
	})
	return
}

func CreateResourceHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"status": true,
		"type":   "ttrn:::api/resources/create",
		"resource": []string{
			"/api/resources/resources",
		},
	})
	return
}

func UpdateResourceHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"type":   "ttrn:::api/resources/update",
		"resource": []string{
			"/api/resources/platforms",
		},
	})
	return
}

func DeleteResourceHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"type":   "ttrn:::api/resources/delete",
		"resource": []string{
			"/api/resources/paas",
		},
	})
	return
}
