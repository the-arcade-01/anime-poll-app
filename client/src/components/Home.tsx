import {useNavigate} from "react-router";

function Home() {
  const navigate = useNavigate();

  return (
    <div>
      <h1>Hello, Anime Poll App</h1>
      <button onClick={() => navigate('/leaderboard')}>Go to Leaderboard</button>
      <button onClick={() => navigate('/fight')}>Go to Fight</button>
    </div>
  );
}

export default Home;