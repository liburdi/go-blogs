#!/bin/bash
docker stop go_blog
docker rm $(docker ps -a -q)
timeNow=$(date "+%Y%m%d%H%M%S")
git pull
go build -o blog
docker build -t golangschool:$timeNow .
docker run -t -i -d -p 127.0.0.1:8080:8080  --name go_blog golangschool:$timeNow