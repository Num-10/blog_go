#!/usr/bin/env bash
#test
echo "111"
docker build -t blog_go -f /home/lnmp/golang/blog_go/Dockerfile .
echo "222"