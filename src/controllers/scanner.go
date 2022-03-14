package controllers

import (
	// built-in
	"encoding/json"
	"net/http"

	// web framework
	"github.com/gin-gonic/gin"

	// local packages
	models "ttcertman/models"
	db "ttcertman/src/platform/database"
	scanner "ttcertman/src/services/scanner"
)

func ScanSSLHandler(c *gin.Context) {
	// validate request
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

	// check permission
	// ...

	// scan
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

	// map to model
	var scan models.Scan
	scan.Name = url
	scan.Key = url
	// convert cert object to json string
	bson, err := json.Marshal(certs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/scan/ssl",
			"type":    "scan-ssl/marshal",
			"message": err.Error(),
		})
		return
	}
	scan.Value = string(bson)

	// create scan on db
	if err := db.GetDbConn().Create(&scan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"intent":  "ttrn:::api/scan/ssl",
			"type":    "scan-ssl/save",
			"message": err,
		})
		return
	}

	// return
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"intent":  "ttrn:::api/scan/ssl",
		"type":    "scan-ssl",
		"message": "SSL certificate scanned successfully",
		"certs":   certs,
	})
	return
}
