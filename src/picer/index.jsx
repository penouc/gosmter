import React, { Component } from "react";

export default class Picer extends Component {
  //   constructor() {
  //     super();
  //   }

  componentDidMount() {}

  render() {
    return (
      <div className="pic-wrap">
        {this.props.src ? (
          <img
            id={this.props.id}
            className="pic-img"
            alt="pic"
            src={this.props.src}
          />
        ) : (
          <img
            id={this.props.id}
            className="pic-img no-pic"
            alt="pic"
            // style={{ display: this.props.src ? "inline" : "none" }}
          />
        )}
      </div>
    );
  }
}
