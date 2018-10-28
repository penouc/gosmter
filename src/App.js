import React, { Component } from "react";
import Picer from "./picer";
import Radios from "./radios";
import Uploader from "./uploader";
import "./App.css";

class App extends Component {
  state = {
    originSrc: "./images/src.jpg",
    targetSrc: "./images/src.jpg"
  };

  changeOriginSrc = url => {
    this.setState({
      originSrc: url,
      targetSrc: url
    });
  };

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <p>
            Gosmter&nbsp;
            <span className="header-sec-title">( Based on Go gift )</span>
          </p>
        </header>
        <div className="main-content">
          <Uploader changeOriginSrc={this.changeOriginSrc.bind(this)} />
          <div className="filter-select">
            <Radios />
          </div>
          <Picer
            id="pic-origin"
            hasLoading={false}
            src={this.state.originSrc}
          />
          <Picer
            id="pic-target"
            hasLoading={true}
            alt="转换后图片"
            src={this.state.targetSrc}
          />
        </div>
      </div>
    );
  }
}

export default App;
