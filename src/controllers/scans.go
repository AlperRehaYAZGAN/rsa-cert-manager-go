/**
*	This is the Scan CRUD Controller.
 */
package controllers

import (
	// built-in
	"net/http"

	// web framework
	"github.com/gin-gonic/gin"
)

func GetScansHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"intent": "ttrn:::api/scans/get",
		"type":   "get-scans",
		"scans": []string{
			"/api/scans/generate/",
		},
	})
	return
}

func CreateScanHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"status": true,
		"intent": "ttrn:::api/scans/create",
		"type":   "create-scan",
		"scan": []string{
			"/api/scan/platforms",
		},
	})
	return
}

func UpdateScanHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"intent": "ttrn:::api/scans/update",
		"type":   "update-scan",
		"scan": []string{
			"/api/scan/platforms",
		},
	})
	return
}

func DeleteScanHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"intent": "ttrn:::api/scans/delete",
		"type":   "delete-scan",
		"scan": []string{
			"/api/scans/paas",
		},
	})
	return
}
