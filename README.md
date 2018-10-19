# gosmter

一个基于 go 编译的 wasm 的实现了图片滤镜的 web 应用

A web Application what used to do images filter , which Based on wasm that written by Go-lang .


## Demo 

访问地址: [https://penouc.com/gosmter/](https://penouc.com/gosmter/)

demo url: [https://penouc.com/gosmter/](https://penouc.com/gosmter/)

## 介绍 / Introduction

这个项目是基于 `create-react-app` 项目上构建的，在 `create-react-app` 的基础上，修改 `webpack` 的版本到 v4， v4 支持 `wasm` 文件的加载，并且修改了 `webpackDev` 的 config , 将加载的 `wasm` 文件的 `MIME Type` 修改为原本的 `application/wasm` , 这样得以保证 `wasm` 在加载为流文件的时候，不会发生错误。

## 依赖 / dependencies
  + go > 1.11 
  + nodejs > 8 
  
  > 如果 go 的版本有问题，执行 `brew upgrade go` ， 或者 `nodejs` 的版本有问题，执行 `nvm install v8.11.0 && nvm use v8.11.0`
  然后需要安装 go 的依赖 `gift` ，
  项目地址 [https://github.com/disintegration/gift](https://github.com/disintegration/gift) 
  执行 `go get -u github.com/disintegration/gift`
  接下来，在主目录下执行 `npm run start` 或者 `yarn start` 即可
  
  
##


