import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from "./pages/Login";
import Admins from "./pages/Admins";
import Users from "./pages/Users";
import Streams from "./pages/Streams";
import Music from "./pages/Music";
import Home from "./pages/Home";

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/home" element={<Home />} />
        <Route path="/admins" element={<Admins />} />
        <Route path="/users" element={<Users />} />
        <Route path="/streams" element={<Streams />} />
        <Route path="/music" element={<Music />} />
      </Routes>
    </Router>
  );
};

export default App;
