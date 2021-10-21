FROM golang:alpine AS builder

# 环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

# 切换到工作目录
WORKDIR /build

# 将代码复制到容器中
COPY . .
# 下载依赖
RUN go mod download

# 编译成二进制文件,二进制文件名：app
RUN go build -o app .


### --------- 二阶段，构建一个小镜像 ---------
FROM debian:stretch-slim

# 项目目录
RUN mkdir www
WORKDIR /www

# 复制配置文件
ARG APP_ENV
COPY ./config-${APP_ENV}.yaml /www/config.yaml

# 从builder镜像中把二进制文件/build/app 拷贝到当前目录
COPY --from=builder /build/app /www

# 修改时区
RUN ln -snf /usr/share/zoneinfo/PRC /etc/localtime && echo PRC > /etc/timezone

