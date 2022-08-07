import {React, useState} from 'react';
import searchImg from './search_grass.png';
import {useNavigate} from 'react-router-dom';

function SearchForm(props) {
  let [state, setState] = useState(props.initState);
  let navigate = useNavigate();

  // とりまnull消すけどこれ辞めないと
  if (state==="null" || state==="undefined" || state===null || state===undefined) {
    setState("")
  }

  function handleChange(event) {
    setState(event.target.value);
  }

  function handleSubmit(event) {
    event.preventDefault();
    if (state !== ""){
      alert(state)
      navigate('../search', {replace:true, state:state});
    } 
  }
  
  return (
    <div className="Search">
      <form onSubmit={handleSubmit} >
        <img src={searchImg} className="Search-img" alt="searchGrass" />
        <label>
          <input type="text" class="Search-input" placeholder="Pipoで検索" value={state} onChange={handleChange} />
        </label>
      </form>
    </div>
  );
}

// propsのデフォルト値
SearchForm.defaultProps = {
  initState: "",
};

export default SearchForm;