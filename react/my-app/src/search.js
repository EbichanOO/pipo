import {React, useState} from 'react';
import searchImg from './search_grass.png';
import {useNavigate} from 'react-router-dom';

function SearchForm() {

  let state = useState();
  let navigate = useNavigate();

  function handleChange(event) {
    state.value = event.target.value;
  }

  function handleSubmit(event) {
      navigate('search');
      alert('A name was submitted: ' + state.value);
      event.preventDefault();
  }
  
  return (
    <div className="Search">
      <form onSubmit={handleSubmit} >
        <img src={searchImg} className="Search-img" alt="searchGrass" />
        <label>
          <input type="text" class="Search-input" placeholder="Pipoで検索" value={state.value} onChange={handleChange} />
        </label>
      </form>
    </div>
  );
}

export default SearchForm;