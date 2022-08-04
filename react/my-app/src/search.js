import searchImg from './search_grass.png';
import React from 'react';

class SearchForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {value: ''};
    
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
      }
    
      handleChange(event) {
        this.setState({value: event.target.value});
      }
    
      handleSubmit(event) {
        alert('A name was submitted: ' + this.state.value);
        event.preventDefault();
      }
    
      render() {
        return (
          <div className="Search">
            <form onSubmit={this.handleSubmit} >
              <img src={searchImg} className="Search-img" alt="searchGrass" />
              <label>
                <input type="text" class="Search-input" placeholder="Pipoで検索" value={this.state.value} onChange={this.handleChange} />
              </label>
            </form>
          </div>
        );
      }
}

export default SearchForm;