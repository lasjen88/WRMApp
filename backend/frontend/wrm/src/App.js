import React, { Component } from 'react';
import logo from './logo.svg';
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
        //mode: 'no-cors',
        headers: {
        "Access-Control-Allow-Origin": "*",
        //"Access-Control-Allow-Origin": "http://localhost:8000/v1/characters",
        "Accept": "application/json",
        //'Content-Type': 'application/json'
      }})
      .then(response => {console.log(response); return response.json()});
      //.then(data => {console.log(data); return data;})
      //.then(data => this.setState({characters: data}))
      //.catch(error => console.log(error));
      return response;
   }
  
  render() {


  this.state.characters.map((item, i) => (
        <div>
          <h1>{ item }</h1>
        </div>

/*
        <div>
          <h1>{ item.name.first }</h1>
          <span>{ item.cell }, { item.email }</span>
        </div>
        */
      ));

    return (

      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h2>Welcome to React</h2>
        </div>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
      </div>
    );
  }
}

export default App;
