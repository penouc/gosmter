import React, { Component } from "react";
import { Icon } from "antd";

export default class Picer extends Component {
  render() {
    return (
      <div className="pic-wrap">
        {this.props.hasLoading ? (
          <Icon
            id="tranform-loading"
            style={{
              display: "none",
              fontSize: "64px",
              color: "#ccc",
              position: "absolute",
              top: "50%",
              left: "50%",
              transform: "translate(-50%, -50%)"
            }}
            type="loading"
            theme="outlined"
          />
        ) : null}
        {this.props.src ? (
          <img
            id={this.props.id}
            className="pic-img"
            alt="pic"
            crossOrigin="Anonymous"
            src={this.props.src}
          />
        ) : (
          <img
            id={this.props.id}
            crossOrigin="Anonymous"
            className="pic-img no-pic"
            alt="pic"
            // style={{ display: this.props.src ? "inline" : "none" }}
          />
        )}
      </div>
    );
  }
}
