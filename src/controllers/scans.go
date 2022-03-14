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
		"type":   "ttrn:::api/scans/get",
		"scans": []string{
			"/api/scans/generate/",
		},
	})
	return
}

func CreateScanHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"status": true,
		"type":   "ttrn:::api/scans/create",
		"scan": []string{
			"/api/scan/platforms",
		},
	})
	return
}

func UpdateScanHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"type":   "ttrn:::api/scans/update",
		"scan": []string{
			"/api/scan/platforms",
		},
	})
	return
}

func DeleteScanHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"type":   "ttrn:::api/scans/delete",
		"scan": []string{
			"/api/scans/paas",
		},
	})
	return
}
