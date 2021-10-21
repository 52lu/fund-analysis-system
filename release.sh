# !/bin/sh
# 编译
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 /usr/local/go/bin/go build -o  app &&
# 运行
docker-compose build  &&
# 启动
docker-compose up -d
