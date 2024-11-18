import React from "react";
import Sidebar from "../components/Sidebar";
import { useNavigate } from "react-router-dom";
import "../layout/Home.css";

const Home = () => {
  const navigate = useNavigate();

  // Handle logout by clearing the token and redirecting to the login page
  const handleLogout = () => {
    localStorage.removeItem("adminToken");
    navigate("/");
  };

  return (
    <div className="home-container">
      <Sidebar />
      <div className="home-content">
        <h1>Welcome to Nexus Music Admin Dashboard</h1>
        <p>You are successfully logged in as an admin.</p>
        <button className="logout-button" onClick={handleLogout}>
          Log Out
        </button>
      </div>
    </div>
  );
};

export default Home;
