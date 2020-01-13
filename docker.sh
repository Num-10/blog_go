#!/usr/bin/env bash
cd /home/lnmp/golang/blog_go
docker build -t blog_go .
echo docker inspect blog_go | grep Id
if docker inspect blog_go | grep Id
then
  docker stop blog_go
  docker rm -f blog_go
  docker run --name blog_go -v /home/lnmp/golang/blog_go/runtime:/home/lnmp/golang/blog_go/runtime --network lnmp -p 8888:8888 -d blog_go