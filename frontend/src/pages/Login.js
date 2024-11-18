import React, { useEffect, useRef, useState } from "react";
import { useNavigate } from "react-router-dom";
import "../layout/Login.css";
import { adminLogin } from "../services/api"; // API service for backend requests

// Import SVG assets for animation
import note1 from "../assets/note1.svg";
import note2 from "../assets/note2.svg";

const Login = () => {
  const canvasRef = useRef(null);
  const navigate = useNavigate();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  // Redirect to the home page if the admin is already logged in
  useEffect(() => {
    const token = localStorage.getItem("adminToken");
    if (token) {
      navigate("/home");
    }
  }, [navigate]);

  // Handle login form submission
  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(""); // Clear previous errors

    try {
      const response = await adminLogin(username, password); // Call admin login API
      localStorage.setItem("adminToken", response.data.token); // Store token in localStorage
      navigate("/home"); // Redirect to home page
    } catch (err) {
      setError("Invalid username or password. Please try again.");
    }
  };

  // Canvas animation logic remains the same as before
  useEffect(() => {
    const canvas = canvasRef.current;
    const ctx = canvas.getContext("2d");

    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;

    const notes = [];
    const svgPaths = [note1, note2];
    const colors = ["#ff512f", "#23a2f6", "#f09819", "#1845ad", "#7C78B8"];

    function random(min, max) {
      return Math.random() * (max - min) + min;
    }

    async function createColoredImage(svgPath, color) {
      const svg = await fetch(svgPath).then((res) => res.text());
      const coloredSvg = svg.replace(/fill="#?[0-9a-fA-F]{3,6}"/g, `fill="${color}"`);

      const img = new Image();
      const svgBlob = new Blob([coloredSvg], { type: "image/svg+xml;charset=utf-8" });
      const url = URL.createObjectURL(svgBlob);

      img.src = url;
      return new Promise((resolve) => {
        img.onload = () => {
          URL.revokeObjectURL(url);
          resolve(img);
        };
      });
    }

    async function createNote(x, y, svgPath) {
      const color = colors[Math.floor(Math.random() * colors.length)];
      const img = await createColoredImage(svgPath, color);

      return {
        x,
        y,
        dx: random(-2, 2),
        dy: random(-2, 2),
        scale: random(1.5, 3),
        angle: random(0, Math.PI * 2),
        angularSpeed: random(0.01, 0.05),
        img,
      };
    }

    function drawNote(note) {
      const width = note.img.width / 8;
      const height = note.img.height / 8;

      ctx.save();
      ctx.translate(note.x, note.y);
      ctx.rotate(note.angle);
      ctx.scale(note.scale, note.scale);
      ctx.globalAlpha = 0.8;

      ctx.drawImage(note.img, -width / 2, -height / 2, width, height);
      ctx.restore();
    }

    function updateNote(note) {
      note.x += note.dx;
      note.y += note.dy;
      note.angle += note.angularSpeed;

      if (note.x + (note.img.width / 8) * note.scale > canvas.width || note.x - (note.img.width / 8) * note.scale < 0) {
        note.dx = -note.dx;
      }
      if (note.y + (note.img.height / 8) * note.scale > canvas.height || note.y - (note.img.height / 8) * note.scale < 0) {
        note.dy = -note.dy;
      }
    }

    async function addNotes() {
      for (let i = 0; i < 10; i++) {
        const note = await createNote(
          random(0, canvas.width),
          random(0, canvas.height),
          svgPaths[Math.floor(Math.random() * svgPaths.length)]
        );
        notes.push(note);
      }
    }

    function animate() {
      ctx.clearRect(0, 0, canvas.width, canvas.height);

      notes.forEach((note) => {
        updateNote(note);
        drawNote(note);
      });

      requestAnimationFrame(animate);
    }

    addNotes().then(() => animate());

    const resizeCanvas = () => {
      canvas.width = window.innerWidth;
      canvas.height = window.innerHeight;
    };
    window.addEventListener("resize", resizeCanvas);

    return () => {
      window.removeEventListener("resize", resizeCanvas);
    };
  }, []);

  return (
    <>
      <canvas ref={canvasRef} className="background"></canvas>
      <form onSubmit={handleSubmit}>
        <h3>Admin Login</h3>

        {error && <p className="error-message">{error}</p>}

        <label htmlFor="username">Username</label>
        <input
          type="text"
          placeholder="Admin Username"
          id="username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
        />

        <label htmlFor="password">Password</label>
        <input
          type="password"
          placeholder="Password"
          id="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />

        <button type="submit">Log In</button>
      </form>
    </>
  );
};

export default Login;
