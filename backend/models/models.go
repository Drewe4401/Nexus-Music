// models/models.go
package models

import "time"

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type Admin struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type Song struct {
	ID       int    `db:"id"`
	Title    string `db:"title"`
	Artist   string `db:"artist"`
	Album    string `db:"album"`
	FilePath string `db:"file_path"` // Path to the audio file
}

type Playlist struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	UserID int    `db:"user_id"` // Foreign key for User
}

type Favorite struct {
	ID     int `db:"id"`
	UserID int `db:"user_id"`
	SongID int `db:"song_id"`
}

type Stream struct {
	ID              int       `db:"id" json:"id"`
	UserID          int       `db:"user_id" json:"user_id"`
	UserUsername    string    `db:"user_username" json:"user_username"` // Fetched from JOIN
	SongID          int       `db:"song_id" json:"song_id"`
	SongTitle       string    `db:"song_title" json:"song_title"` // Fetched from JOIN
	StreamedAt      time.Time `db:"streamed_at" json:"streamed_at"`
	DurationSeconds int       `db:"duration_seconds" json:"duration_seconds"`
}
