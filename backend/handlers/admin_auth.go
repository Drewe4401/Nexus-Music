package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"nexus-music/db"
	"nexus-music/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var adminJwtKey = []byte("admin_secret_key")

// AdminLoginRequest represents the admin login payload
type AdminLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AdminLogin handles admin authentication
func AdminLogin(c *gin.Context) {
	var req AdminLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Fetch the admin from the database
	var hashedPassword string
	var adminID int
	err := db.DB.QueryRow("SELECT id, password FROM admins WHERE username = $1", req.Username).Scan(&adminID, &hashedPassword)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	// Compare the provided password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate a JWT token for the admin
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"adminID": adminID,
		"exp":     jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})
	tokenString, err := token.SignedString(adminJwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func AuthenticateAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return adminJwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Extract adminID from the token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		adminID, ok := claims["adminID"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Store adminID in the context
		c.Set("adminID", int(adminID))
		c.Next()
	}
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
