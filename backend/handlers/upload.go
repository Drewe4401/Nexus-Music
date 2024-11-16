// handlers/upload.go
package handlers

import (
	"fmt"
	"net/http"
	"nexus-music/db"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadMusic(c *gin.Context) {
	userID, _ := c.Get("userID") // Get user ID from JWT claims

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}

	// Define the path to save the file (using user ID to organize)
	savePath := filepath.Join("uploads", fmt.Sprintf("%d", userID), file.Filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Insert song record into database
	_, err = db.DB.Exec("INSERT INTO songs (user_id, title, file_path) VALUES ($1, $2, $3)", userID, file.Filename, savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
