#镜像
FROM ubuntu:18.04
#安装工具
RUN apt-get update -y -q && apt-get upgrade -y -q \
    && apt-get install gcc -y -q  \
    && gcc --version    \
    && apt-get install wget -y -q 
#安装go,清理go文件
RUN wget https://studygolang.com/dl/golang/go1.15.4.linux-amd64.tar.gz \
    && tar xfz go1.15.4.linux-amd64.tar.gz -C /usr/local    \
    && rm -f go1.15.4.linux-amd64.tar.gz
#设置环境变量
ENV PATH $PATH:/usr/local/go/bin
ENV GO111MODULE on
ENV GOPROXY=https://goproxy.cn,direct
RUN go version
#设置工作目录
WORKDIR $GOPATH/src/github.com/KouKouChan/CSO2-Server
#下载mod
COPY go.mod .
COPY go.sum .
RUN go mod download
#复制项目文件
COPY . .
#清理项目git
RUN rm -rf ./.git
#构建项目
RUN GOOS=linux GOARCH=amd64 go build -o CSO2-Server-docker .
#设置工作目录
WORKDIR $GOPATH/src/github.com/KouKouChan/
#切换可执行文件位置
RUN mv ./CSO2-Server/CSO2-Server-docker /usr/local/CSO2-Server-docker
RUN mv ./CSO2-Server /usr/local/CSO2-Server
#设置工作目录
WORKDIR /usr/local/
#清理项目
RUN rm -rf /usr/local/go
RUN rm -rf /root/go
RUN rm -rf /root/.cache
RUN apt-get remove gcc -q -y
RUN apt-get remove wget -q -y
#暴露端口
EXPOSE 1314
EXPOSE 1315
EXPOSE 30001
EXPOSE 30002
#最终运行docker的命令
#USER app-runner
ENTRYPOINT  ["./CSO2-Server-docker"]