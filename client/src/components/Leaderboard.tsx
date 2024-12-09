import {useEffect, useState} from "react";
import {AnimeDetails} from "./types.ts";
import {useNavigate} from "react-router";

function Leaderboard() {
  const navigate = useNavigate();
  const [leaderboard, setLeaderboard] = useState<AnimeDetails[]>([]);

  const fetchLeaderBoard = async () => {
    try {
      const res = await fetch(`${import.meta.env.VITE_API_URL}/leaderboard/10`);
      const data = await res.json();
      console.log(data);
      setLeaderboard(data);
    } catch (error) {
      console.error("Error fetching anime data:", error);
    }
  }

  useEffect(() => {
    fetchLeaderBoard().catch(error => console.error('Error in useEffect:', error));
  }, []);

  return (
    <div>
      <h2>Leaderboard</h2>
      <button onClick={() => navigate('/')}>Go to Home</button>
      <ul>
        {leaderboard.map((anime: AnimeDetails, index) => (
          <li key={index} style={{listStyleType: 'none', marginBottom: '20px'}}>
            <img src={anime.image_link} alt={anime.title} style={{width: '100px', height: '150px'}}/>
            <p>{anime.title}</p>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Leaderboard;