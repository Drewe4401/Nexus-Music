import axios from "axios";

// Create an Axios instance for the API
const api = axios.create({
  baseURL: process.env.REACT_APP_API_URL || "http://localhost:8080", // Use environment variable or default to localhost
});

// Add a request interceptor to include the token in every request
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem("adminToken"); // Retrieve the admin token
    if (token) {
      config.headers.Authorization = `Bearer ${token}`; // Add token to headers
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Admin authentication
export const adminLogin = (username, password) => {
  return api.post("/admin/login", { username, password });
};

// Fetch all users
export const fetchUsers = () => {
  return api.get("/admin/users");
};

// Fetch all streams
export const fetchStreams = () => {
  return api.get("/admin/streams");
};

// Fetch all music
export const fetchMusic = () => {
  return api.get("/admin/music");
};


// Create a new user
export const createUser = (username, password) => {
  return api.post("/admin/users", { username, password });
};

// Update a user's password
export const updateUserPassword = (userId, password) => {
  return api.post("/admin/users/update-password", { user_id: userId, password });
};

// Delete a user
export const deleteUser = (userId) => {
  return api.delete(`/admin/users/${userId}`);
};

// Fetch all admins
export const fetchAdmins = () => {
  return api.get("/admin/admins");
};

// Create a new admin
export const createAdmin = (username, password) => {
  return api.post("/admin/admins/create-admin", { username, password });
};

// Update an admin's password
export const updateAdminPassword = (adminId, password) => {
  return api.post("/admin/admins/update-password", { admin_id: adminId, password });
};

// Delete an admin
export const deleteAdmin = (adminId) => {
  return api.delete(`/admin/admins/${adminId}`);
};


export default api;
