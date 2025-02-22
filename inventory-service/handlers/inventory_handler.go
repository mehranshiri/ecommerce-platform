package handlers

import (
	"inventory-service/database"
	"inventory-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateInventory(c *gin.Context) {
	var input models.Inventory

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "bad request",
			"detail": err.Error(),
		})
		return
	}

	if err := models.Create(database.DB, &input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "bad request",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "created successfully",
		"inventory": input,
	})
}
