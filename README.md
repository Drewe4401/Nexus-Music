# 🎵 Nexus Music

**Nexus Music** is a self-hosted music streaming platform that empowers users to manage and stream their music collections. The application includes a backend powered by Go, an admin frontend built with React, and a mobile app coded in React Native for streaming on the go. All components are containerized with Docker for seamless deployment.

---

## 🚀 Features

- **Admin Dashboard**: Manage users, upload music, and analyze streaming activity.
- **Music Streaming**: Stream your music library securely from anywhere.
- **Multi-Platform**:
  - Admin Frontend: Built in React for managing users and music.
  - Mobile App: React Native application for end users to stream music.
- **Analytics**: Track and visualize user streaming activity.
- **Self-Hosting**: Fully containerized using Docker for easy setup and deployment.

---

## 📂 Project Structure

```plaintext
nexus-music/
├── backend/               # Go backend code
│   ├── db/                # Database migrations and utilities
│   ├── handlers/          # API handlers for authentication, streaming, etc.
│   ├── models/            # Data models for database entities
│   ├── main.go            # Entry point for the Go backend
│   └── go.mod             # Go modules file
│
├── admin-panel/           # React admin frontend
│   ├── public/            # Static files
│   ├── src/               # Source files for the React app
│   ├── package.json       # Dependencies for the frontend
│   ├── Dockerfile         # Dockerfile for the admin app
│   └── .env               # Environment variables for the frontend
│
├── docker-compose.yml     # Docker Compose configuration for backend and frontend
├── README.md              # Documentation for the project
└── .gitignore             # Ignored files for the project
