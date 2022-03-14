/**
*	This is the Resource CRUD Controller.
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

func GetResourcesHandler(ctx *gin.Context) {
	// validate request
	var pagination models.Pagination
	if err := ctx.ShouldBindJSON(&pagination); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/resources/get",
			"type":    "get-resources/validation",
			"message": err.Error(),
		})
		return
	}

	// check permission
	// ...

	// get resources
	var resources []models.Resource
	if err := db.GetDbConn().Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&resources); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/resources/get",
			"type":    "get-resources/select",
			"message": err,
		})
		return
	}

	// return
	ctx.JSON(http.StatusOK, gin.H{
		"status":    true,
		"intent":    "ttrn:::api/resources/get",
		"type":      "get-resources",
		"resources": resources,
	})
	return
}

func CreateResourceHandler(ctx *gin.Context) {
	// validate request
	var createResourceDto models.CreateResourceDto
	if err := ctx.ShouldBindQuery(&createResourceDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/resources/create",
			"type":    "create-resource/validation",
			"message": err.Error(),
		})
		return
	}

	// check permission
	// ...

	// map to model
	resource := models.Resource{
		Name: createResourceDto.Name,
		Key:  createResourceDto.Key,
	}

	// create resource
	if err := db.GetDbConn().Create(&resource); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/resources/create",
			"type":    "create-resource/save",
			"message": err,
		})
		return
	}

	// return
	ctx.JSON(http.StatusCreated, gin.H{
		"status":   true,
		"intent":   "ttrn:::api/resources/create",
		"type":     "create-resource",
		"resource": resource,
	})
	return
}

func UpdateResourceHandler(ctx *gin.Context) {
	// validate request
	var updateResourceDto models.UpdateResourceDto
	if err := ctx.ShouldBindJSON(&updateResourceDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/resources/update",
			"type":    "update-resource/validation",
			"message": err.Error(),
		})
		return
	}

	// check permission
	// ...

	// map to model
	resource := models.Resource{
		Name: updateResourceDto.Name,
		Key:  updateResourceDto.Key,
	}

	// update resource
	if err := db.GetDbConn().Create(&resource); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/resources/update",
			"type":    "update-resource/save",
			"message": err,
		})
		return
	}

	// return
	ctx.JSON(http.StatusOK, gin.H{
		"status":   true,
		"intent":   "ttrn:::api/resources/update",
		"type":     "update-resource",
		"resource": resource,
	})
	return
}

func DeleteResourceHandler(ctx *gin.Context) {
	// validate request
	var deleteResourceDto models.DeleteResourceDto
	if err := ctx.ShouldBindJSON(&deleteResourceDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/resources/delete",
			"type":    "delete-resource/validation",
			"message": err.Error(),
		})
		return
	}

	// check permission
	// ...

	// delete resource
	if err := db.GetDbConn().Delete(&models.Resource{}, "key = ?", deleteResourceDto.Key); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/resources/delete",
			"type":    "delete-resource/save",
			"message": err,
		})
		return
	}

	// return
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
