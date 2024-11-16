// handlers/search.go
package handlers

import (
	"net/http"
	"nexus-music/db"

	"github.com/gin-gonic/gin"
)

func SearchMusic(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter required"})
		return
	}

	// Query songs by title or artist
	rows, err := db.DB.Query("SELECT id, title, artist FROM songs WHERE title ILIKE $1 OR artist ILIKE $1", "%"+query+"%")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var id int
		var title, artist string
		if err := rows.Scan(&id, &title, &artist); err == nil {
			results = append(results, map[string]interface{}{
				"id":     id,
				"title":  title,
				"artist": artist,
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}
