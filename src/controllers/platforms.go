/**
*	This is the Platform CRUD Controller.
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

func GetPlatformsHandler(ctx *gin.Context) {
	// validate request
	var pagination models.Pagination
	// if limit and page is null set defaults
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/platforms/get",
			"type":    "get-platforms/validation",
			"message": err.Error(),
		})
		return
	}

	// check permission
	// ...

	// select from db
	var platforms []models.Platform
	db.GetDbConn().Limit(pagination.Limit).Offset(pagination.Limit * (pagination.Page - 1)).Find(&platforms)

	// return success
	ctx.JSON(http.StatusOK, gin.H{
		"status":    true,
		"intent":    "ttrn:::api/platforms/get",
		"type":      "get-platforms",
		"platforms": platforms,
	})
	return
}

func CreatePlatformHandler(ctx *gin.Context) {
	// validate request
	var createPlatformDto models.CreatePlatformDto
	if err := ctx.ShouldBindJSON(&createPlatformDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/platforms/create",
			"type":    "create-platform/validation",
			"message": err.Error(),
		})
		return
	}

	// check permission
	// ...

	// map to model
	platform := models.Platform{
		Name: createPlatformDto.Name,
		Key:  createPlatformDto.Key,
	}

	// create platform
	if db.GetDbConn().Create(&platform).Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"intent": "ttrn:::api/platforms/create",
			"type":   "create-platform/save",
		})
		return
	}

	// return success
	ctx.JSON(http.StatusCreated, gin.H{
		"status":   true,
		"intent":   "ttrn:::api/platforms/create",
		"type":     "create-platform",
		"message":  "Platform created successfully",
		"platform": platform,
	})
	return
}

func UpdatePlatformHandler(ctx *gin.Context) {
	// validate request
	var updatePlatformDto models.UpdatePlatformDto
	if err := ctx.ShouldBindJSON(&updatePlatformDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/platforms/update",
			"type":    "update-platform/validation",
			"message": err.Error(),
		})
		return
	}

	// check permission
	// ...

	// map to model
	platform := models.Platform{
		Name: updatePlatformDto.Name,
		Key:  updatePlatformDto.Key,
	}

	// update platform
	if err := db.GetDbConn().Save(&platform); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/platforms/update",
			"type":    "update-platform/save",
			"message": err,
		})
		return
	}

	// return success
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
	// validate request
	var deletePlatformDto models.DeletePlatformDto
	if err := ctx.ShouldBindJSON(&deletePlatformDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/platforms/delete",
			"type":    "delete-platform/validation",
			"message": err.Error(),
		})
		return
	}

	// check permission
	// ...

	// map to model
	platform := models.Platform{
		Key: deletePlatformDto.Key,
	}

	// delete platform
	if err := db.GetDbConn().Delete(&platform); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/platforms/delete",
			"type":    "delete-platform/save",
			"message": err,
		})
		return
	}

	// return success
	ctx.JSON(http.StatusOK, gin.H{
		"status":   true,
		"intent":   "ttrn:::api/platforms/delete",
		"type":     "delete-platform",
		"platform": platform,
	})
	return
}
