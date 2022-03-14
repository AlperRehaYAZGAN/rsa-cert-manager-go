package controllers

import (
	// built-in
	"net/http"

	// local packages
	scanner "ttcertman/src/services/scanner"

	// web framework
	"github.com/gin-gonic/gin"
)

func ScanSSLHandler(c *gin.Context) {
	url := c.DefaultQuery("url", "google.com:443")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "url is required",
		})
		return
	}

	certs, err := scanner.GetCertificates(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "scan ssl certificate",
		"certs":   certs,
	})
	return
}
