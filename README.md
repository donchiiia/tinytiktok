# TinyTiktok
## 项目介绍
- 字节跳动第五届青训营后端项目
- 本项目基于 kitex RPC微服务 + Hertz HTTP服务实现
- IDL使用 `thirft`
- API层使用了HTTP协议， Service层使用了RPC协议
- 使用ETCD作为注册发现的组件
- 使用MySQL作为数据库
- 使用Gorm作为ORM操作组件
- 使用Minio实现OSS存储
- 使用Jaeger实现链路追踪
- 使用JWT实现用户token鉴权

## 项目成员
队友组队失败，最终本人单独完成。。。

## 参考文档
本项目基于青训营提供的[抖音项目方案说明](https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof#K6ToR1)实现接口开发，并通过[极简抖音App](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)完成效果测试

## 模块介绍
|Service Name| Usage | Framework | protocol | Path | IDL |
|:---:|:---:|:---:|:---:|:---:|:---:|
|api|HTTP interface|kitex/hertz|http|cmd/api|idl/api.thrift|
|user|user management|kitex|thrift|cmd/user|idl/user.thrift|
|feed|feed management|kitex|thrift|cmd/feed|idl/feed.thrift|
|relation|relation management|kitex|thrift|cmd/relation|idl/relation.thrift|
|favorite|favorite management|kitex|thrift|cmd/favorite|idl/favorite.thrift|
|comment|comment management|kitex|thrift|cmd/comment|idl/comment.thrift|
|publish|publish management|kitex|thrift|cmd/publish|idl/publish.thrift|
|message|message management|kitex|thrift|cmd/message|idl/message.thrift|

## 架构图
![架构图](/images/架构图.png)

## ER图
![ER diagram](/images/ER.png)

## 项目目录介绍
```
.
├── README.md
├── cmd // 服务包
│   ├── api
│   ├── comment
│   ├── favorite
│   ├── feed
│   ├── message
│   ├── publish
│   ├── relation
│   └── user
├── config // nginx配置
│   └── nginx.conf
├── dal    // 数据访问层
│   ├── db
│   └── init.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── idl    // idl文件夹
│   ├── api.thrift
│   ├── comment.thrift
│   ├── favorite.thrift
│   ├── feed.thrift
│   ├── message.thrift
│   ├── publish.thrift
│   ├── relation.thrift
│   └── user.thrift
├── images
├── kitex_gen
├── pack
├── pkg
│   ├── configs
│   ├── consts // 全局变量
│   ├── errno  // 自定义错误
│   ├── ffmpeg // ffmpeg调用
│   ├── jwt    // jwt鉴权
│   ├── minio  // minio操作
│   └── mw     // 通用日志中间件
└── test
    └── bear.mp4
```
## 快速开始
**启动前请修改`pkg\consts\consts.go`中的配置项，以便符合你的设备**
- Setup Basic Dependence
```
docker-compose up
```
- Run User RPC Server
```
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```
- Run Feed RPC Server
```
cd cmd/feed
sh build.sh
sh output/bootstrap.sh
```
- Run Publish RPC Server
```
cd cmd/publish
sh build.sh
sh output/bootstrap.sh
```
- Run Favorite RPC Server
```
cd cmd/favorite
sh build.sh
sh output/bootstrap.sh
```
- Run Comment RPC Server
```
cd cmd/comment
sh build.sh
sh output/bootstrap.sh
```
- Run Relation RPC Server
```
cd cmd/relation
sh build.sh
sh output/bootstrap.sh
```
- Run Message RPC Server
```
cd cmd/message
sh build.sh
sh output/bootstrap.sh
```
- Run API Server
```
cd cmd/api
go run .
```