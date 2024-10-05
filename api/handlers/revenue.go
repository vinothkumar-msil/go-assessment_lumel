package handlers

import (
	"github.com/gin-gonic/gin"
	"go-backend-assessment/internal/calculations"
	"net/http"
)

// RevenueHandler calculates and returns the total revenue for the specified date range
func RevenueHandler(c *gin.Context) {
	startDate := c.Query("start") // Get the 'start' query parameter
	endDate := c.Query("end")     // Get the 'end' query parameter

	// Call the revenue calculation logic
	revenue, err := calculations.CalculateRevenue(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_revenue": revenue,
	})
}

