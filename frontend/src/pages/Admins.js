import React, { useEffect, useState } from "react";
import Sidebar from "../components/Sidebar";
import {
  fetchAdmins,
  createAdmin,
  updateAdminPassword,
  deleteAdmin,
} from "../services/api";
import "../layout/Admins.css";

const Admins = () => {
  const [admins, setAdmins] = useState([]);
  const [newAdmin, setNewAdmin] = useState({ username: "", password: "" });
  const [selectedAdmin, setSelectedAdmin] = useState(null);
  const [showCreateForm, setShowCreateForm] = useState(false); // Toggle for "Create Admin" form
  const [notification, setNotification] = useState(""); // Notification message

  useEffect(() => {
    fetchAllAdmins();
  }, []);

  const fetchAllAdmins = async () => {
    try {
      const response = await fetchAdmins();
      setAdmins(response.data.admins);
    } catch (error) {
      console.error("Failed to fetch admins", error);
    }
  };

  const handleDelete = async (id) => {
    if (window.confirm("Are you sure you want to delete this admin?")) {
      try {
        await deleteAdmin(id);
        fetchAllAdmins();
        setNotification("Admin deleted successfully!");
        setTimeout(() => setNotification(""), 3000);
      } catch (error) {
        console.error("Failed to delete admin", error);
      }
    }
  };

  const handleUpdatePassword = async () => {
    try {
      await updateAdminPassword(selectedAdmin.id, selectedAdmin.newPassword);
      fetchAllAdmins();
      setSelectedAdmin(null); // Close the form
      setNotification("Admin password updated successfully!");
      setTimeout(() => setNotification(""), 3000);
    } catch (error) {
      console.error("Failed to update password", error);
    }
  };

  const handleCreateAdmin = async () => {
    try {
      await createAdmin(newAdmin.username, newAdmin.password);
      fetchAllAdmins();
      setNewAdmin({ username: "", password: "" });
      setShowCreateForm(false); // Close the form
      setNotification("Admin created successfully!");
      setTimeout(() => setNotification(""), 3000);
    } catch (error) {
      console.error("Failed to create admin", error);
    }
  };

  return (
    <div className="main-container">
      <Sidebar />
      <div className="main-content">
        <h1>Admins Page</h1>
        <p>Manage all administrators here.</p>

        {/* Notification */}
        {notification && <div className="notification">{notification}</div>}

        {/* Create New Admin Button */}
        <div className="header-actions">
          <button
            className="create-button"
            onClick={() => setShowCreateForm(!showCreateForm)}
          >
            Create New Admin
          </button>
        </div>

        {/* Admin Table */}
        <div className="table-container">
          <table>
            <thead>
              <tr>
                <th>ID</th>
                <th>Username</th>
                <th>Actions</th>
              </tr>
            </thead>
            <tbody>
              {admins && admins.length > 0 ? (
                admins.map((admin) => (
                  <tr key={admin.ID}>
                    <td>{admin.ID}</td>
                    <td>{admin.Username}</td>
                    <td>
                      <button
                        onClick={() =>
                          setSelectedAdmin({
                            id: admin.ID,
                            username: admin.Username,
                            newPassword: "",
                          })
                        }
                      >
                        Change Password
                      </button>
                      <button onClick={() => handleDelete(admin.ID)}>
                        Delete
                      </button>
                    </td>
                  </tr>
                ))
              ) : (
                <tr>
                  <td colSpan="3" style={{ textAlign: "center" }}>
                    No admins found.
                  </td>
                </tr>
              )}
            </tbody>
          </table>
        </div>

        {/* Change Password Form */}
        {selectedAdmin && (
          <div className="form-container">
            <button
              className="close-button"
              onClick={() => setSelectedAdmin(null)}
            >
              X
            </button>
            <form className="form">
              <h3>Change Password for {selectedAdmin.username}</h3>
              <input
                type="password"
                placeholder="New Password"
                value={selectedAdmin.newPassword}
                onChange={(e) =>
                  setSelectedAdmin({
                    ...selectedAdmin,
                    newPassword: e.target.value,
                  })
                }
              />
              <button onClick={handleUpdatePassword} type="button">
                Update Password
              </button>
            </form>
          </div>
        )}

        {/* Create Admin Form */}
        {showCreateForm && (
          <div className="form-container">
            <button
              className="close-button"
              onClick={() => setShowCreateForm(false)}
            >
              X
            </button>
            <form className="form">
              <h3>Create New Admin</h3>
              <input
                type="text"
                placeholder="Username"
                value={newAdmin.username}
                onChange={(e) =>
                  setNewAdmin({ ...newAdmin, username: e.target.value })
                }
              />
              <input
                type="password"
                placeholder="Password"
                value={newAdmin.password}
                onChange={(e) =>
                  setNewAdmin({ ...newAdmin, password: e.target.value })
                }
              />
              <button onClick={handleCreateAdmin} type="button">
                Create Admin
              </button>
            </form>
          </div>
        )}
      </div>
    </div>
  );
};

export default Admins;
