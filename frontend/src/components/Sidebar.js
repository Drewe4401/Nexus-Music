// src/components/Sidebar.js
import React from "react";
import { NavLink } from "react-router-dom";
import "../layout/Sidebar.css";

const Sidebar = () => {
  return (
    <div className="sidebar">
      <div className="sidebar-header">
        <h2>Nexus Music</h2>
      </div>
      <ul className="sidebar-list">
        <li className="sidebar-list-item">
          <NavLink 
            to="/home" 
            className={({ isActive }) => (isActive ? "active" : "")}
          >
            Home
          </NavLink>
        </li>
        <li className="sidebar-list-item">
          <NavLink 
            to="/admins" 
            className={({ isActive }) => (isActive ? "active" : "")}
          >
            Admins
          </NavLink>
        </li>
        <li className="sidebar-list-item">
          <NavLink 
            to="/users" 
            className={({ isActive }) => (isActive ? "active" : "")}
          >
            Users
          </NavLink>
        </li>
        <li className="sidebar-list-item">
          <NavLink 
            to="/streams" 
            className={({ isActive }) => (isActive ? "active" : "")}
          >
            Streams
          </NavLink>
        </li>
        <li className="sidebar-list-item">
          <NavLink 
            to="/music" 
            className={({ isActive }) => (isActive ? "active" : "")}
          >
            Music
          </NavLink>
        </li>
      </ul>
    </div>
  );
};

export default Sidebar;
