package handlers

import (
	"net/http"
	"nexus-music/db"
	"nexus-music/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

// GetAllMusic retrieves all songs from the database
func GetAllMusic(c *gin.Context) {
	var songs []models.Song
	err := db.DB.Select(&songs, "SELECT id, title, artist, album, file_path FROM songs")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve music"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"songs": songs})
}

// GetAllAdmins retrieves all admins from the database
func GetAllAdmins(c *gin.Context) {
	var admins []models.Admin
	err := db.DB.Select(&admins, "SELECT id, username FROM admins")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve admins"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"admins": admins})
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

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	_, err = db.DB.Exec("UPDATE users SET password = $1 WHERE id = $2", string(hashedPassword), request.UserID)
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

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	_, err = db.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// CreateAdmin creates a new admin in the database
func CreateAdmin(c *gin.Context) {
	var request struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if the username already exists
	var existingCount int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM admins WHERE username = $1", request.Username).Scan(&existingCount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing username"})
		return
	}

	if existingCount > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Insert the new admin into the database
	_, err = db.DB.Exec("INSERT INTO admins (username, password) VALUES ($1, $2)", request.Username, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin created successfully"})
}

func UpdateAdminPassword(c *gin.Context) {
	var request struct {
		AdminID  int    `json:"admin_id"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	_, err = db.DB.Exec("UPDATE admins SET password = $1 WHERE id = $2", string(hashedPassword), request.AdminID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin password updated successfully"})
}

// DeleteAdmin deletes an admin from the database
func DeleteAdmin(c *gin.Context) {
	adminID := c.Param("id")

	_, err := db.DB.Exec("DELETE FROM admins WHERE id = $1", adminID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete admin"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin deleted successfully"})
}
