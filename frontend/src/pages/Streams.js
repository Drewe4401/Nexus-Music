import React, { useEffect, useState } from "react";
import Sidebar from "../components/Sidebar";
import { fetchStreams } from "../services/api";
import '../layout/Streams.css';

const Streams = () => {
  const [streams, setStreams] = useState([]);

  useEffect(() => {
    fetchAllStreams();
  }, []);

  const fetchAllStreams = async () => {
    try {
      const response = await fetchStreams();
      setStreams(response.data.streams);
    } catch (error) {
      console.error("Failed to fetch streams", error);
    }
  };

  return (
    <div className="main-container">
      <Sidebar />
      <div className="main-content">
        <h1>Streams Page</h1>
        <p>View all song streams here.</p>

        <div className="table-container">
          <table>
            <thead>
              <tr>
                <th>ID</th>
                <th>Username</th>
                <th>Song Title</th>
                <th>Streamed At</th>
                <th>Duration (seconds)</th>
              </tr>
            </thead>
            <tbody>
              {streams && streams.length > 0 ? (
                streams.map((stream) => (
                  <tr key={stream.id}>
                    <td>{stream.id}</td>
                    <td>{stream.user_username}</td>
                    <td>{stream.song_title}</td>
                    <td>{new Date(stream.streamed_at).toLocaleString()}</td>
                    <td>{stream.duration_seconds}</td>
                  </tr>
                ))
              ) : (
                <tr>
                  <td colSpan="5" style={{ textAlign: "center" }}>
                    No streams found.
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

export default Streams;
