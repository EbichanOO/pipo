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
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
      </header>
      <body className="App-body">
        <SearchForm />
      </body>
    </div>
  );
}

export default App;
