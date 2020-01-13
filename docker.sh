#!/usr/bin/env bash
cd /home/lnmp/golang/blog_go
docker build -t blog_go .

if [ "$(docker ps -q -f name=blog_go)" ]; then
    if [ "$(docker ps -aq -f status=running -f name=blog_go)" ]; then
        # stop
        docker stop blog_go
    fi
    if [ "$(docker ps -aq -f status=exited -f name=blog_go)" ]; then
        # cleanup
        docker rm blog_go
    fi
fi
# run container
docker run --name blog_go -v /home/lnmp/golang/blog_go/runtime:/home/lnmp/golang/blog_go/runtime --network lnmp -p 8888:8888 -d blog_go
