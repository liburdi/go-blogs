### 前言

2018年接触golang后，非常喜欢此语言的特性和风格。

遂搭建[golang中文社区](https://www.golangschool.com)，为了记录自己的学习的点点滴滴，也为了和更多人一起交流学习。

### 概览

- go version 

```1.12以上```

- 框架

无/原生

- go mod 

```
require (
	github.com/garyburd/redigo v1.6.2
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xorm/xorm v0.7.9
	xorm.io/builder v0.3.7 // indirect
	xorm.io/core v0.7.3 // indirect
)
```
- 目录结构

1.common 公用的助手函数

2.config 配置文件

3.controllers 控制器，存放api文件

4.db 数据库连接，数据库方法重写等

5.redis redis连接，redis方法重写等

6.models 数据库模型

7.router 路由

8.templates 模版文件目录

9.static 静态文件目录

10.uploads 上传文件目录


### 下载
git clone https://github.com/KenRitchie/golangschool.git

### 启动

```
#默认
cd golangschool
go run main.go

#docker

cd golangschool
./build.sh

```
别忘了根据golangschool.sql创建表结构，以及修改config对应的配置文件





