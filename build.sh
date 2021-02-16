#!/bin/bash

if [ $1 = "stop" ]; then
  docker stop go_blog
fi

if [ $1 =  "restart" ]; then
  docker stop go_blog
  ./build.sh start
fi

if [ $1 =  "start" ]; then
  docker rm $(docker ps -a -q)
  timeNow=$(date "+%Y%m%d%H%M%S")

  git pull
  if [ ! -d static ]; then
          mkdir static
  fi
  if [ ! -d uploads ]; then
          mkdir uploads
  fi
  go build -o golangschool
  docker build  -t golangschool:$timeNow .
  docker run -t -i -d -p 8083:8083  --name go_blog golangschool:$timeNow

fi