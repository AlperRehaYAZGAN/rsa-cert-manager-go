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
			"status":  false,
			"intent":  "ttrn:::api/scan/ssl",
			"type":    "scan-ssl",
			"message": "URL is required",
		})
		return
	}

	certs, err := scanner.GetCertificates(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/scan/ssl",
			"type":    "scan-ssl",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"intent":  "ttrn:::api/scan/ssl",
		"type":    "scan-ssl",
		"message": "SSL certificate scanned successfully",
		"certs":   certs,
	})
	return
}
