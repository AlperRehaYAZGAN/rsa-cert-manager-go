/**
*	This is the Scan CRUD Controller.
 */
package controllers

import (
	// built-in
	"net/http"

	// web framework
	"github.com/gin-gonic/gin"

	// local packages
	models "ttcertman/models"
	db "ttcertman/src/platform/database"
)

func GetScansHandler(ctx *gin.Context) {
	// validate request
	var pagination models.Pagination
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/scans/get",
			"type":    "get-scans/validation",
			"message": err.Error(),
		})
		return
	}

	// check permission
	// ...

	// get scans
	var scans []models.Scan
	if err := db.GetDbConn().Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&scans); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/scans/get",
			"type":    "get-scans/select",
			"message": err,
		})
		return
	}

	// return
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"intent": "ttrn:::api/scans/get",
		"type":   "get-scans",
		"scans":  scans,
	})
	return
}

func DeleteScanHandler(ctx *gin.Context) {
	// validate request
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/scans/delete",
			"type":    "delete-scan/validation",
			"message": "ID is required",
		})
		return
	}

	// check permission
	// ...

	// delete scan
	if err := db.GetDbConn().Delete(&models.Scan{}, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/scans/delete",
			"type":    "delete-scan/save",
			"message": err,
		})
		return
	}

	// return
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"intent": "ttrn:::api/scans/delete",
		"type":   "delete-scan",
		"scan":   id,
	})
	return
}
