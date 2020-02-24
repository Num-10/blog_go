#源镜像
FROM golang:latest
#作者
MAINTAINER num_10 "num_10@163.com"
#设置工作目录
WORKDIR /home/lnmp/golang/blog_go
#将服务器的go工程代码加入到docker容器中
ADD . /home/lnmp/golang/blog_go
#go构建可执行文件
RUN GO111MODULE=on GOPROXY="https://gocenter.io" go build .
#暴露端口
EXPOSE 8888
#最终运行docker的命令
ENTRYPOINT  ["./blog_go"]