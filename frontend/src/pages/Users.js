import React, { useEffect, useState } from "react";
import Sidebar from "../components/Sidebar";
import { fetchUsers, createUser, updateUserPassword, deleteUser } from "../services/api";
import '../layout/Users.css';

const Users = () => {
  const [users, setUsers] = useState([]);
  const [newUser, setNewUser] = useState({ username: "", password: "" });
  const [selectedUser, setSelectedUser] = useState(null);
  const [showCreateForm, setShowCreateForm] = useState(false); // State to toggle form visibility
  const [notification, setNotification] = useState(""); // For success notification

  useEffect(() => {
    fetchAllUsers();
  }, []);

  const fetchAllUsers = async () => {
    try {
      const response = await fetchUsers();
      setUsers(response.data.users);
    } catch (error) {
      console.error("Failed to fetch users", error);
    }
  };
  

  const handleDelete = async (id) => {
    if (window.confirm("Are you sure you want to delete this user?")) {
      try {
        await deleteUser(id);
        fetchAllUsers();
      } catch (error) {
        console.error("Failed to delete user", error);
      }
    }
  };

  const handleUpdatePassword = async () => {
    try {
      await updateUserPassword(selectedUser.id, selectedUser.newPassword);
      fetchAllUsers();
      setSelectedUser(null);
    } catch (error) {
      console.error("Failed to update password", error);
    }
  };

  const handleCreateUser = async () => {
    try {
      await createUser(newUser.username, newUser.password);
      fetchAllUsers();
      setNewUser({ username: "", password: "" });
      setShowCreateForm(false); // Hide form after successful creation
      setNotification("User created successfully!"); // Set notification
      setTimeout(() => setNotification(""), 3000); // Clear notification after 3 seconds
    } catch (error) {
      console.error("Failed to create user", error);
    }
  };

  return (
    <div className="main-container">
      <Sidebar />
      <div className="main-content">
        <h1>Users Page</h1>
        <p>Manage all app users here.</p>

        {/* Notification */}
        {notification && <div className="notification">{notification}</div>}

        {/* Create New User Button */}
        <div className="header-actions">
          <button
            className="create-button"
            onClick={() => setShowCreateForm(!showCreateForm)}
          >
            Create New User
          </button>
        </div>

        <div className="table-container">

        {showCreateForm && (
  <div>
    <form className="create-form">
    <button
      className="close-button"
      onClick={() => setShowCreateForm(false)}
    >
      X
    </button>
      <h3>Create New User</h3>
      <input
        type="text"
        placeholder="Username"
        value={newUser.username}
        onChange={(e) => setNewUser({ ...newUser, username: e.target.value })}
      />
      <input
        type="password"
        placeholder="Password"
        value={newUser.password}
        onChange={(e) => setNewUser({ ...newUser, password: e.target.value })}
      />
      <button onClick={handleCreateUser} type="button">
        Create User
      </button>
    </form>
  </div>
)}

  <table>
    <thead>
      <tr>
        <th>ID</th>
        <th>Username</th>
        <th></th>
      </tr>
    </thead>
    <tbody>
      {users && users.length > 0 ? (
        users.map((user) => (
          <tr key={user.ID}>
            <td>{user.ID}</td>
            <td>{user.Username}</td>
            <td className="table_button">
              <button
                onClick={() => setSelectedUser({ id: user.ID, username: user.Username, newPassword: "" })}
              >
                Change Password
              </button>
              <button onClick={() => handleDelete(user.ID)}>Delete</button>
            </td>
          </tr>
        ))
      ) : (
        <tr>
          <td colSpan="3" style={{ textAlign: "center" }}>
            No users found.
          </td>
        </tr>
      )}
    </tbody>
  </table>
</div>


        {selectedUser && (
          <form>
                <button
                  className="close-button"
                  onClick={() => setSelectedUser(false)}
               >
                X
                </button>
            <h3>Change {selectedUser.username}'s Password</h3>
            <input
              type="password"
              placeholder="New Password"
              value={selectedUser.newPassword}
              onChange={(e) =>
                setSelectedUser({ ...selectedUser, newPassword: e.target.value })
              }
            />
            <button onClick={handleUpdatePassword}>Update Password</button>
          </form>
        )}

        {/* Conditional rendering of the Create New User form */}
      </div>
    </div>
  );
};

export default Users;
