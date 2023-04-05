import { BrowserRouter as Router, Route, Routes } from "react-router-dom";

import Organisations from "./pages/Organisations";
import Organisation from "./pages/Organisation";

import './App.css';

// redirect hoem page to organisations pg
function App() {
  return (
    <div>
      <Router>
        <Routes>
        <Route path="/" element={<Organisations />} />
        <Route path="/organisations" element={<Organisations />} />
        <Route path="/organisation" element={<Organisation />} />
        </Routes>
      </Router>
    
    </div>
  );
}


export default App;

