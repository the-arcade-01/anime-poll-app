import {BrowserRouter, Route, Routes} from "react-router";
import Home from "./components/Home.tsx";
import Leaderboard from "./components/Leaderboard.tsx";
import Fight from "./components/Fight.tsx";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home/>}/>
        <Route path="/leaderboard" element={<Leaderboard/>}/>
        <Route path="/fight" element={<Fight/>}/>
      </Routes>
    </BrowserRouter>
  );
}

export default App;