import React, { Component } from 'react';
import logo from './wrm.svg';
import './App.css';

class App extends Component {

    constructor(props) {
      super(props);
  
      this.state = {characters: []};
    }
  
    componentDidMount() {
      this.getCharacters();
    }
  
    async getCharacters() {
      let response = await fetch("http://localhost:8000/v1/characters", 
      {
        method: 'GET',
        dataType : 'jsonp',
        headers: {
        "Accept": "application/json",
      }})
      .then(response => {console.log(response); return response.json()})
      .then(data => {console.log(data); return data;})
      .then(data => this.setState({characters: data}))
      .catch(error => console.log(error));
      return response;
   }
  
  render() {


  const characters = this.state.characters.map((character, i) => (
        <div key={character.charactername}>
          <h1> Character: { character.charactername }</h1>
          <span>Warrior: { character.warrior }, 
                Rogue: { character.rogue }, 
                Mage: { character.mage }
          </span>
        </div>
      ));

    return (

      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h2>Welcome to Warrior Rogue Mage</h2>
        </div>
        <div id="layout-content" className="layout-content-wrapper">
          <div className="panel-list">{ characters }</div>
        </div>
      </div>
    );
  }
}

export default App;
