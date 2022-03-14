/**
*	This is the Services CRUD Controller.
 */
package controllers

import (
	// built-in
	"net/http"

	// web framework
	"github.com/gin-gonic/gin"
)

func GetServicesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"intent": "ttrn:::api/services/get",
		"type":   "get-services",
		"services": []string{
			"/api/services/generate/ssh",
		},
	})
	return
}

func CreateServiceHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"status": true,
		"intent": "ttrn:::api/services/create",
		"type":   "create-service",
		"service": []string{
			"/api/services/service",
		},
	})
	return
}

func UpdateServiceHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"intent": "ttrn:::api/services/update",
		"type":   "update-service",
		"service": []string{
			"/api/services/service",
		},
	})
	return
}

func DeleteServiceHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"intent": "ttrn:::api/services/delete",
		"type":   "delete-service",
		"service": []string{
			"/api/services/service",
		},
	})
	return
}
