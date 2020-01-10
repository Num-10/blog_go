#!/usr/bin/env bash
echo "111"
cd /home/lnmp/golang/blog_go
docker stop blog_go
docker rm blog_go
docker build -t blog_go .
echo "222"