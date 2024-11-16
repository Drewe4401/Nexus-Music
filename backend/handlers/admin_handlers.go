package handlers

import (
	"net/http"
	"nexus-music/db"
	"nexus-music/models"

	"github.com/gin-gonic/gin"
)

// GetAllUsers retrieves all users from the database
func GetAllUsers(c *gin.Context) {
	var users []models.User
	err := db.DB.Select(&users, "SELECT id, username FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
