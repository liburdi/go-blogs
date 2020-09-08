#!/bin/bash
timeNow=$(date "+%Y%m%d%H%M%S")
git pull
go build
docker build -t golangschool:$timeNow .
docker run -t -i -d -p 127.0.0.1:8080:8080 golangschool:$timeNow
