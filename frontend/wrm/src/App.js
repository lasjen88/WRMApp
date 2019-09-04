import React, { Component } from 'react';
import logo from './wrm.svg';
import './App.css';
import SideNav, {MenuIcon} from 'react-simple-sidenav';

class Spells extends React.Component {
  constructor (props) {
    super();
    this.state = {
      spells: props.spells
    }
  }
  render (){

  const spells = this.state.spells.map((spell, i) => (
          <div key={spell.spellName}>
            <Spell spell={spell} />
          </div>
        ));

    return (
      <div align="left">
        {spells}
      </div>
    )
  }
}
 

class Spell extends React.Component {
  constructor (props) {
    super()
    this.state = {
      spell: props.spell,
      isHidden: true
    }
  }
  toggleHidden () {
    this.setState({
      isHidden: !this.state.isHidden
    })
  }
  render () {
    return (
      <div>
        <button onClick={this.toggleHidden.bind(this)} >
          {this.state.spell.spellName}
        </button>
        {!this.state.isHidden && <SpellInfo data={this.state.spell}/>}
      </div>
    )
  }
}

function SpellInfo(props){
  let data = props.data;
  return (
    <div align="left">
      <ul>
        <li>Name: {data.spellName}</li>
        <li>Spell Cost: {data.spellCost}</li>
        <li>Level: {data.difficultyLevel}</li>
        <li>Mana Consumption: {data.manaConsumption}</li>
        <li>Description: {data.spellDescription}</li>
      </ul>
    </div>
  )
}


class App extends Component {

    constructor(props) {
      super(props);
      this.URL = 'http://localhost:8000/v1';
      this.state = {
        characters: [],
        spells: ['hello'],
        downloading: true,
      };
    }
  
    async componentDidMount() {
      await this.getCharacters();
      let spells = await this.getSpells();
      this.setState({spells: spells});
      this.setState({downloading: false});
    }

  async getSpells(){
    let data = await fetch(this.URL + "/spells", 
        {
          method: 'GET',
          dataType : 'jsonp',
          headers: {
          "Accept": "application/json",
        }})
        .then(response => {return response.json()})
        .catch(error => console.log(error));
        return data;
  }

  async getCharacters() {
      let response = await fetch(this.URL + "/characters", 
      {
        method: 'GET',
        dataType : 'jsonp',
        headers: {
        "Accept": "application/json",
      }})
      .then(response => {return response.json()})
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

    if (this.state.downloading === true){
      return null;
    }

      return (
        <div className="App">
          <div className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
            <h2>Welcome to Warrior Rogue Mage</h2>
          </div>

          <h2>Spells</h2>
        <MenuIcon onClick={() => this.setState({showNav: true})}/>

        <SideNav
          showNav={this.state.showNav}
          onHideNav={() => this.setState({showNav: false})} />
          <Spells spells={this.state.spells}  />
          
          <div className="panel-list">{ characters }</div>

       </div>
      );
    }
}

export default App;
