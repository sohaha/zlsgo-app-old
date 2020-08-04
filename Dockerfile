FROM golang:alpine as builder

# 设置go mod proxy 国内代理
# 设置golang path
ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    GO111MODULE=on \
    CGO_ENABLED=1
WORKDIR /zlsapp
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
COPY . .
RUN go env && go list && go build -o app main.go

EXPOSE 3788
ENTRYPOINT /zlsapp/app

# 根据 Dockerfile 生成 Docker 镜像
# docker build -t zlsapp .

# 根据 Docker 镜像启动 Docker 容器
# docker run -itd -p 8888:8888 --name zlsapp zlsapp