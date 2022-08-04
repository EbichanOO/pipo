import logo from './logo.svg';
import './App.css';
import SearchForm from './search';

function App() {
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
