#!/bin/bash
#cd /home/lnmp/golang/blog_go
echo '111'
#生成镜像
docker build -t blog_go -f /home/lnmp/golang/blog_go/Dockerfile .
echo '222'
#停止并删除旧容器
#docker stop blog_go
#docker rm blog_go
#创建容器
#docker run --name blog_go -v /home/lnmp/golang/blog_go/runtime:/home/lnmp/golang/blog_go/runtime --network lnmp -p 8888:8888 -d blog_go