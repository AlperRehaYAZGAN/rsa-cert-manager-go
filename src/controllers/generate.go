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
			"status":  false,
			"intent":  "ttrn:::generate/ssl",
			"type":    "generate-ssl",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"intent":  "ttrn:::api/generate/ssl",
		"type":    "generate-ssl",
		"message": "SSL Key pair generated successfully",
		"cert":    cert,
	})
	return
}

func GenerateSSHHandler(ctx *gin.Context) {
	cert, err := certificates.GenerateSSHKeyPair()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/generate/ssh",
			"type":    "generate-ssh",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"intent":  "ttrn:::api/generate/ssh",
		"type":    "generate-ssh",
		"message": "SSH Key pair generated successfully",
		"cert":    cert,
	})
	return
}
