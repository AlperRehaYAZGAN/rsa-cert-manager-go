package controllers

import (
	// built-in
	"net/http"

	// local packages
	certificates "ttcertman/src/services/certificates"

	// web framework
	"github.com/gin-gonic/gin"
)

func GenerateSSLHandler(ctx *gin.Context) {
	cert, err := certificates.GenerateSSLKeyPair()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "generate ssl key pair",
		"cert":    cert,
	})
	return
}

func GenerateSSHHandler(ctx *gin.Context) {
	cert, err := certificates.GenerateSSHKeyPair()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "generate ssh key pair",
		"cert":    cert,
	})
	return
}
