package db

import (
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// RunMigrations ensures that the necessary tables exist in the database
func RunMigrations() {
	migrations := map[string]string{
		"users": `
			CREATE TABLE IF NOT EXISTS users (
				id SERIAL PRIMARY KEY,
				username VARCHAR(255) UNIQUE NOT NULL,
				password VARCHAR(255) NOT NULL
			);`,
		"admins": `
			CREATE TABLE IF NOT EXISTS admins (
				id SERIAL PRIMARY KEY,
				username VARCHAR(255) UNIQUE NOT NULL,
				password VARCHAR(255) NOT NULL
			);`,
		"songs": `
			CREATE TABLE IF NOT EXISTS songs (
				id SERIAL PRIMARY KEY,
				title VARCHAR(255) NOT NULL,
				artist VARCHAR(255) NOT NULL,
				album VARCHAR(255),
				file_path VARCHAR(255) NOT NULL
			);`,
		"playlists": `
			CREATE TABLE IF NOT EXISTS playlists (
				id SERIAL PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				user_id INT NOT NULL,
				FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
			);`,
		"favorites": `
			CREATE TABLE IF NOT EXISTS favorites (
				id SERIAL PRIMARY KEY,
				user_id INT NOT NULL,
				song_id INT NOT NULL,
				FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
				FOREIGN KEY (song_id) REFERENCES songs(id) ON DELETE CASCADE
			);`,
		"streams": `
			CREATE TABLE IF NOT EXISTS streams (
				id SERIAL PRIMARY KEY,
				user_id INT NOT NULL,
				song_id INT NOT NULL,
				streamed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				duration_seconds INT,
				FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
				FOREIGN KEY (song_id) REFERENCES songs(id) ON DELETE CASCADE
			);`,
	}

	// Execute all table creation queries
	for table, query := range migrations {
		_, err := DB.Exec(query)
		if err != nil {
			log.Printf("❌ Error creating table '%s': %v\n", table, err)
			continue
		}
		log.Printf("✅ Table '%s' is ready (created or already existed).\n", table)
	}
	log.Println("✅ All migrations completed.")

	// Ensure default admin user exists
	ensureDefaultAdmin()
}

// ensureDefaultAdmin checks if a default admin exists and creates one if it doesn't
func ensureDefaultAdmin() {
	// Fetch admin credentials from environment variables
	defaultAdminUsername := getEnv("DEFAULT_ADMIN_USERNAME", "admin")
	defaultAdminPassword := getEnv("DEFAULT_ADMIN_PASSWORD", "password")

	// Check if an admin already exists
	var adminCount int
	err := DB.QueryRow("SELECT COUNT(*) FROM admins").Scan(&adminCount)
	if err != nil {
		log.Fatalf("❌ Error checking admin existence: %v", err)
	}

	// If no admin exists, create the default admin
	if adminCount == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(defaultAdminPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("❌ Error hashing default admin password: %v", err)
		}

		_, err = DB.Exec("INSERT INTO admins (username, password) VALUES ($1, $2)", defaultAdminUsername, string(hashedPassword))
		if err != nil {
			log.Fatalf("❌ Error creating default admin: %v", err)
		}

		log.Printf("✅ Default admin created: username=%s", defaultAdminUsername)
	} else {
		log.Println("✅ Admin user already exists. Skipping default admin creation.")
	}
}

// getEnv retrieves the value of an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
