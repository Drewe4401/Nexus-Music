import React, { useEffect, useState } from "react";
import Sidebar from "../components/Sidebar";
import { fetchMusic } from "../services/api"; // Import the API function to fetch music
import "../layout/Music.css"; // CSS for styling

const Music = () => {
  const [music, setMusic] = useState([]); // State to hold the music data

  // Fetch all music on component mount
  useEffect(() => {
    fetchAllMusic();
  }, []);

  const fetchAllMusic = async () => {
    try {
      const response = await fetchMusic();
      setMusic(response.data.songs); // Assume the API response has a `songs` key
    } catch (error) {
      console.error("Failed to fetch music", error);
    }
  };

  return (
    <div className="main-container">
      <Sidebar />
      <div className="main-content">
        <h1>Music Page</h1>
        <p>View all the songs available in the library.</p>

        <div className="table-container">
          <table>
            <thead>
              <tr>
                <th>ID</th>
                <th>Title</th>
                <th>Artist</th>
                <th>Album</th>
                <th>File Path</th>
              </tr>
            </thead>
            <tbody>
              {music.length > 0 ? (
                music.map((song) => (
                  <tr key={song.ID}>
                    <td>{song.ID}</td>
                    <td>{song.Title}</td>
                    <td>{song.Artist}</td>
                    <td>{song.Album || "N/A"}</td>
                    <td>{song.FilePath}</td>
                  </tr>
                ))
              ) : (
                <tr>
                  <td colSpan="5" style={{ textAlign: "center" }}>
                    No music found.
                  </td>
                </tr>
              )}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default Music;
