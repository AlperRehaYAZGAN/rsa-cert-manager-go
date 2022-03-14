package controllers

import (
	// built-in
	"net/http"

	// web framework
	"github.com/gin-gonic/gin"

	// local packages
	models "ttcertman/models"
	db "ttcertman/src/platform/database"
	certificates "ttcertman/src/services/certificates"
)

func GenerateSSLHandler(ctx *gin.Context) {
	// check permission
	// ...

	// generate ssl
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

	// map to model
	var sslKeyPair models.SSLPair
	sslKeyPair.PublicKey = string(cert.PublicKey)
	sslKeyPair.PrivateKey = string(cert.PrivateKey)

	// save to database
	if err := db.GetDbConn().Create(&sslKeyPair).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"intent":  "ttrn:::generate/ssl",
			"type":    "generate-ssl",
			"message": err.Error(),
		})
		return
	}

	// return
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
	// check permission
	// ...

	// generate ssh
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

	// map to model
	var sshKeyPair models.SSHPair
	sshKeyPair.PublicKey = string(cert.PublicKey)
	sshKeyPair.PrivateKey = string(cert.PrivateKey)

	// save to database
	if err := db.GetDbConn().Create(&sshKeyPair).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/generate/ssh",
			"type":    "generate-ssh",
			"message": err.Error(),
		})
		return
	}

	// return
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"intent":  "ttrn:::api/generate/ssh",
		"type":    "generate-ssh",
		"message": "SSH Key pair generated successfully",
		"cert":    cert,
	})
	return
}
