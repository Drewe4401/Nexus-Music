// handlers/playlist.go
package handlers

import (
	"net/http"
	"nexus-music/db"

	"github.com/gin-gonic/gin"
)

type CreatePlaylistRequest struct {
	Name string `json:"name" binding:"required"`
}

func CreatePlaylist(c *gin.Context) {
	userID, _ := c.Get("userID") // Get user ID from JWT claims

	var req CreatePlaylistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Insert the new playlist into the database
	_, err := db.DB.Exec("INSERT INTO playlists (user_id, name) VALUES ($1, $2)", userID, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Playlist created successfully"})
}
