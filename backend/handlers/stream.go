package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"nexus-music/db"

	"github.com/gin-gonic/gin"
)

// getEnv retrieves the value of an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// StreamAudio streams the requested audio file and logs the stream for the user
func StreamAudio(c *gin.Context) {
	songID := c.Param("id")
	userID := c.GetInt("userID") // Get the user ID from the authenticated context

	// Build the path to the music file
	musicDir := getEnv("MUSIC_DIR", "music")
	songPath := filepath.Join(musicDir, songID+".mp3")

	// Open the music file
	file, err := os.Open(songPath)
	if err != nil {
		c.String(http.StatusNotFound, "File not found")
		return
	}
	defer file.Close()

	// Get file information
	fi, _ := file.Stat()
	fileSize := fi.Size()

	// Set response headers
	c.Header("Content-Type", "audio/mpeg")
	c.Header("Accept-Ranges", "bytes")

	// Handle range requests
	rangeHeader := c.GetHeader("Range")
	if rangeHeader == "" {
		// Serve the entire file if no range is specified
		c.File(songPath)
		logStream(userID, songID, 0) // Log stream with 0 duration for full file
		return
	}

	// Parse the range header
	rangeParts := rangeHeader[6:]
	start, _ := strconv.ParseInt(rangeParts, 10, 64)
	end := fileSize - 1
	if dashIndex := len(rangeParts); dashIndex > 0 {
		if parsedEnd, err := strconv.ParseInt(rangeParts[dashIndex+1:], 10, 64); err == nil {
			end = parsedEnd
		}
	}

	// Serve the requested byte range
	c.Status(http.StatusPartialContent)
	c.Header("Content-Range", "bytes "+strconv.FormatInt(start, 10)+"-"+strconv.FormatInt(end, 10)+"/"+strconv.FormatInt(fileSize, 10))
	c.Header("Content-Length", strconv.FormatInt(end-start+1, 10))

	sectionReader := io.NewSectionReader(file, start, end-start+1)
	http.ServeContent(c.Writer, c.Request, file.Name(), fi.ModTime(), sectionReader)

	// Log the stream after serving content
	duration := int((end - start) / 1000) // Convert bytes to approximate seconds
	logStream(userID, songID, duration)
}

// logStream logs a song stream in the database
func logStream(userID int, songID string, duration int) {
	_, err := db.DB.Exec(`
		INSERT INTO streams (user_id, song_id, streamed_at, duration_seconds)
		VALUES ($1, $2, $3, $4)`,
		userID, songID, time.Now(), duration,
	)
	if err != nil {
		// Log the error, but don't stop the response
		// as logging the stream is non-critical
		// (i.e., the user still gets to listen to their music)
		println("Error logging stream:", err.Error())
	}
}
