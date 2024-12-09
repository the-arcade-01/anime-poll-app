import {useEffect, useState} from "react";
import {AnimeDetails} from "./types.ts";
import {useNavigate} from "react-router";

function Fight() {
  const navigate = useNavigate();
  const [animes, setAnimes] = useState<AnimeDetails[]>([]);

  const fetchAnimesForFaceOff = async () => {
    try {
      const res = await fetch(`${import.meta.env.VITE_API_URL}/fight`);
      const data = await res.json();
      console.log(data);
      setAnimes(data);
    } catch (error) {
      console.error("Error fetching anime data:", error);
    }
  }

  const submitVote = async (mal_id: number) => {
    try {
      const response = await fetch(`${import.meta.env.VITE_API_URL}/vote/${mal_id}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        }
      });
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      console.log('Vote submitted successfully');
    } catch (error) {
      console.error('Error on submitting vote:', error);
    } finally {
      window.location.reload();
    }
  };

  useEffect(() => {
    fetchAnimesForFaceOff().catch(error => console.error('Error in useEffect:', error));
  }, []);

  return (
    <div style={{display: 'flex', justifyContent: 'space-around'}}>
      <button onClick={() => navigate('/')}>Go to Home</button>
      {animes.map((anime: AnimeDetails, index) => (
        <div key={index} style={{textAlign: 'center', cursor: "pointer"}} onClick={() => submitVote(anime.mal_id)}>
          <img src={anime.image_link} alt={anime.title} style={{width: '200px', height: '300px'}}/>
          <p>{anime.title}</p>
        </div>
      ))}
    </div>
  );
}

export default Fight;