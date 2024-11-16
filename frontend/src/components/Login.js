import React, { useEffect, useRef } from "react";
import "../layout/Login.css";

// Import the SVG assets as strings for manipulation
import note1 from "../assets/note1.svg";
import note2 from "../assets/note2.svg";

const Login = () => {
  const canvasRef = useRef(null);

  useEffect(() => {
    const canvas = canvasRef.current;
    const ctx = canvas.getContext("2d");

    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;

    const notes = []; // Array to store note objects
    const svgPaths = [note1, note2]; // SVG file paths
    const colors = ["#ff512f", "#23a2f6", "#f09819", "#1845ad", "#7C78B8"]; // Available colors

    // Function to create random values
    function random(min, max) {
      return Math.random() * (max - min) + min;
    }

    // Function to create a colored SVG image
    async function createColoredImage(svgPath, color) {
      const svg = await fetch(svgPath).then((res) => res.text());
      const coloredSvg = svg.replace(/fill="#?[0-9a-fA-F]{3,6}"/g, `fill="${color}"`);

      const img = new Image();
      const svgBlob = new Blob([coloredSvg], { type: "image/svg+xml;charset=utf-8" });
      const url = URL.createObjectURL(svgBlob);

      img.src = url;
      return new Promise((resolve) => {
        img.onload = () => {
          URL.revokeObjectURL(url); // Cleanup blob
          resolve(img);
        };
      });
    }

    // Create a note object
    async function createNote(x, y, svgPath) {
      const color = colors[Math.floor(Math.random() * colors.length)];
      const img = await createColoredImage(svgPath, color);

      return {
        x,
        y,
        dx: random(-2, 2),
        dy: random(-2, 2),
        scale: random(1.5, 3), // Scale down the image by a factor of 8
        angle: random(0, Math.PI * 2), // Random rotation
        angularSpeed: random(0.01, 0.05), // Speed of rotation
        img,
      };
    }

    // Draw the note on the canvas
    function drawNote(note) {
      const width = note.img.width / 8; // Adjust width
      const height = note.img.height / 8; // Adjust height

      ctx.save();
      ctx.translate(note.x, note.y);
      ctx.rotate(note.angle);
      ctx.scale(note.scale, note.scale);
      ctx.globalAlpha = 0.8; // Add transparency

      ctx.drawImage(note.img, -width / 2, -height / 2, width, height);
      ctx.restore();
    }

    // Update the note's position and rotation
    function updateNote(note) {
      note.x += note.dx;
      note.y += note.dy;
      note.angle += note.angularSpeed;

      // Bounce off walls
      if (note.x + (note.img.width / 8) * note.scale > canvas.width || note.x - (note.img.width / 8) * note.scale < 0) {
        note.dx = -note.dx;
      }
      if (note.y + (note.img.height / 8) * note.scale > canvas.height || note.y - (note.img.height / 8) * note.scale < 0) {
        note.dy = -note.dy;
      }
    }

    // Add random notes to the array
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

    // Animate the notes
    function animate() {
      ctx.clearRect(0, 0, canvas.width, canvas.height);

      notes.forEach((note) => {
        updateNote(note);
        drawNote(note);
      });

      requestAnimationFrame(animate);
    }

    addNotes().then(() => animate());

    // Handle window resize
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
      <form>
        <h3>Login Here</h3>

        <label htmlFor="username">Username</label>
        <input type="text" placeholder="Admin Username" id="username" required />

        <label htmlFor="password">Password</label>
        <input type="password" placeholder="Password" id="password" required />

        <button type="submit">Log In</button>
      </form>
    </>
  );
};

export default Login;
