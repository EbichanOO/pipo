import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import logo from './logo.svg';
import './App.css';
import SearchForm from './search';
import InfoPage from './informationPage';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="" element={<Home />} />
        <Route path="search" element={<InfoPage />} />
      </Routes>
    </Router>
  );
}

function Home() {
  return (
    <div className="Index">
      <header className="Index-header">
        <img src={logo} className="Index-logo" alt="logo" />
      </header>
      <body className="Index-body">
        <SearchForm />
      </body>
    </div>
  );
}

export default App;
