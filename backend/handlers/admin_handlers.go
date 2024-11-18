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

// GetAllStreams retrieves all streams from the database
func GetAllStreams(c *gin.Context) {
	var streams []models.Stream
	query := `
		SELECT 
			st.id, st.user_id, st.song_id, st.streamed_at, st.duration_seconds,
			u.username AS user_username, 
			s.title AS song_title 
		FROM streams st
		JOIN users u ON st.user_id = u.id
		JOIN songs s ON st.song_id = s.id
		ORDER BY st.streamed_at DESC
	`
	err := db.DB.Select(&streams, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve streams"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"streams": streams})
}

func UpdateUserPassword(c *gin.Context) {
	var request struct {
		UserID   int    `json:"user_id"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.DB.Exec("UPDATE users SET password = $1 WHERE id = $2", request.Password, request.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

// DeleteUser deletes a user from the database
func DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	_, err := db.DB.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// CreateUser creates a new user in the database
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
