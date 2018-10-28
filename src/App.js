import React, { Component } from "react";
import Picer from "./picer";
import Radios from "./radios";
import "./App.css";

class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <p>Gosmter</p>
        </header>
        <div className="main-content">
          <div className="filter-select">
            <Radios />
          </div>
          <Picer id="pic-origin" src="./images/src.jpg" />
          <Picer id="pic-target" />
        </div>
      </div>
    );
  }
}

export default App;
