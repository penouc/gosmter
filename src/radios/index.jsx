import React, { Component } from "react";
import { Radio } from "antd";

const RadioGroup = Radio.Group;

export default class Radios extends Component {
  state = {
    value: "0"
  };

  onChange = e => {
    window.transforImg2Base(e.target.value);
    this.setState({
      value: e.target.value
    });
  };

  render() {
    return (
      <RadioGroup
        name="radiogroup"
        onChange={this.onChange}
        value={this.state.value}
      >
        <Radio value={"0"}>无</Radio>
        <Radio value={"resize"}>resize</Radio>
        <Radio value={"crop_to_size"}>crop_to_size</Radio>
        <Radio value={"rotate_180"}>rotate_180</Radio>
        <Radio value={"rotate_30"}>rotate_30</Radio>
        <Radio value={"brightness_increase"}>brightness_increase</Radio>{" "}
        <Radio value={"brightness_decrease"}>brightness_decrease</Radio>{" "}
        <Radio value={"contrast_increase"}>contrast_increase</Radio>
        <Radio value={"contrast_decrease"}>contrast_decrease</Radio>
        <Radio value={"saturation_increase"}>saturation_increase</Radio>{" "}
        <Radio value={"saturation_decrease"}>saturation_decrease</Radio>{" "}
        <Radio value={"gamma_1.5"}>gamma_1.5</Radio>
        <Radio value={"gamma_0.5"}>gamma_0.5</Radio>
        <Radio value={"gaussian_blur"}>gaussian_blur</Radio>
        <Radio value={"unsharp_mask"}>unsharp_mask</Radio>
        <Radio value={"sigmoid"}>sigmoid</Radio>
        <Radio value={"pixelate"}>pixelate</Radio>
        <Radio value={"colorize"}>colorize</Radio>
        <Radio value={"grayscale"}>grayscale</Radio>
        <Radio value={"sepia"}>sepia</Radio>
        <Radio value={"invert"}>invert</Radio>
        <Radio value={"mean"}>mean</Radio>
        <Radio value={"median"}>median</Radio>
        <Radio value={"minimum"}>minimum</Radio>
        <Radio value={"maximum"}>maximum</Radio>
        <Radio value={"hue_rotate"}>hue_rotate</Radio>
        <Radio value={"color_balance"}>color_balance</Radio>
      </RadioGroup>
    );
  }
}
