/**
*	This is the Services CRUD Controller.
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

func GetServicesHandler(ctx *gin.Context) {
	// validate request
	var pagination models.Pagination
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/services/get",
			"type":    "get-services/validation",
			"message": err.Error(),
		})
		return
	}

	// check permission
	// ...

	// get services
	var services []models.Service
	if err := db.GetDbConn().Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&services); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/services/get",
			"type":    "get-services/select",
			"message": err,
		})
		return
	}

	// return
	ctx.JSON(http.StatusOK, gin.H{
		"status":   true,
		"intent":   "ttrn:::api/services/get",
		"type":     "get-services",
		"services": services,
	})
	return
}

func CreateServiceHandler(ctx *gin.Context) {
	// validate request
	var createServiceDto models.CreateServiceDto
	if err := ctx.ShouldBindJSON(&createServiceDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/services/create",
			"type":    "create-service/validation",
			"message": err.Error(),
		})
		return
	}

	// check permission
	// ...

	// map to model
	service := models.Service{
		Name: createServiceDto.Name,
		Key:  createServiceDto.Key,
	}

	// create service
	if err := db.GetDbConn().Create(&service); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/services/create",
			"type":    "create-service/save",
			"message": err,
		})
		return
	}

	// return
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"intent":  "ttrn:::api/services/create",
		"type":    "create-service",
		"service": service,
	})
	return
}

func UpdateServiceHandler(ctx *gin.Context) {
	// validate request
	var updateServiceDto models.UpdateServiceDto
	if err := ctx.ShouldBindJSON(&updateServiceDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/services/update",
			"type":    "update-service/validation",
			"message": err.Error(),
		})
		return
	}

	// check permission
	// ...

	// map to model
	service := models.Service{
		Name: updateServiceDto.Name,
		Key:  updateServiceDto.Key,
	}

	// update service
	if err := db.GetDbConn().Save(&service); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/services/update",
			"type":    "update-service/save",
			"message": err,
		})
		return
	}

	// return
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"intent":  "ttrn:::api/services/update",
		"type":    "update-service",
		"service": service,
	})
	return
}

func DeleteServiceHandler(ctx *gin.Context) {
	// validate request
	var deleteServiceDto models.DeleteServiceDto
	if err := ctx.ShouldBindJSON(&deleteServiceDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/services/delete",
			"type":    "delete-service/validation",
			"message": err.Error(),
		})
		return
	}

	// check permission
	// ...

	// map to model
	service := models.Service{
		Name: deleteServiceDto.Name,
		Key:  deleteServiceDto.Key,
	}

	// delete service
	if err := db.GetDbConn().Delete(&service); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/services/delete",
			"type":    "delete-service/delete",
			"message": err,
		})
		return
	}

	// return
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"intent":  "ttrn:::api/services/delete",
		"type":    "delete-service",
		"service": service,
	})
	return
}
