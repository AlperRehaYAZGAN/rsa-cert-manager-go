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
		"intent": "ttrn:::api/resources/get",
		"type":   "get-resources",
		"resources": []string{
			"/api/services/generate/",
		},
	})
	return
}

func CreateResourceHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"status": true,
		"intent": "ttrn:::api/resources/create",
		"type":   "create-resource",
		"resource": []string{
			"/api/resources/resources",
		},
	})
	return
}

func UpdateResourceHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"intent": "ttrn:::api/resources/update",
		"type":   "update-resource",
		"resource": []string{
			"/api/resources/platforms",
		},
	})
	return
}

func DeleteResourceHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"intent": "ttrn:::api/resources/delete",
		"type":   "delete-resource",
		"resource": []string{
			"/api/resources/paas",
		},
	})
	return
}
