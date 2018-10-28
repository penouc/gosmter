import React, { Component } from "react";
import { Upload, Button, Icon } from "antd";

import "antd/dist/antd.css";

export default class Uploader extends Component {
  state = {
    fileList: []
  };

  render() {
    let _this = this;
    this.changeOriginSrc = this.props.changeOriginSrc;
    const props = {
      action: "https://sm.ms/api/upload",
      onRemove: file => {
        this.setState(({ fileList }) => {
          const index = fileList.indexOf(file);
          const newFileList = fileList.slice();
          newFileList.splice(index, 1);
          return {
            fileList: newFileList
          };
        });
      },
      beforeUpload: file => {
        var reader = new FileReader();
        reader.onload = function(e) {
          // target.result 该属性表示目标对象的DataURL
          _this.changeOriginSrc(e.target.result);
        };
        // 传入一个参数对象即可得到基于该参数对象的文本内容
        reader.readAsDataURL(file);
        return false;
      },
      fileList: this.state.fileList
    };

    return (
      <div className="uploader-buttons">
        <Upload {...props}>
          <Button>
            <Icon type="upload" /> 选择图片文件
          </Button>
        </Upload>
      </div>
    );
  }
}
