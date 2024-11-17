import React from "react";
import { useNavigate } from "react-router-dom";

const Home = () => {
  const navigate = useNavigate();

  // Handle logout by clearing the token and redirecting to the login page
  const handleLogout = () => {
    localStorage.removeItem("adminToken");
    navigate("/");
  };

  return (
    <div style={styles.container}>
      <h1>Welcome to Nexus Music Admin Dashboard</h1>
      <p>You are successfully logged in as an admin.</p>
      <button style={styles.button} onClick={handleLogout}>
        Log Out
      </button>
    </div>
  );
};

// Simple inline styles for the page
const styles = {
  container: {
    textAlign: "center",
    padding: "50px",
    fontFamily: "Arial, sans-serif",
  },
  button: {
    padding: "10px 20px",
    backgroundColor: "#ff512f",
    color: "white",
    border: "none",
    borderRadius: "5px",
    cursor: "pointer",
    fontSize: "16px",
  },
};

export default Home;
