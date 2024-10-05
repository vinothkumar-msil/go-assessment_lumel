package handlers

import (
	"github.com/gin-gonic/gin"
	"go-backend-assessment/scripts"
	"net/http"
)

// RefreshDataHandler is the handler for triggering data refresh via API
func RefreshDataHandler(c *gin.Context) {
	err := scripts.LoadCSVData("data.csv") // Path to your CSV file
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error refreshing data",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data refresh initiated successfully.",
	})
}
