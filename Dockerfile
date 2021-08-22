# 1. 编译 Go 程序
# 设置Go编译器版本，默认为最新版
ARG GO_VERSION=latest

FROM golang:${GO_VERSION} AS builder

# 镜像元数据，可以使用 docker inspect 查看
LABEL author="SsrCoder@gmail.com"

# 设置工作目录
WORKDIR $GOPATH/src/github.com/SsrCoder/leetwatcher

# 将当前文件夹下的文件拷贝到工作目录
COPY . .

# 静态链接编译Go程序
RUN GO111MODULE=on GOPROXY="https://goproxy.cn,direct" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o output/output .

####################################################
# 2. 压缩
FROM ssrcoder/upx as compress

COPY --from=builder /go/src/github.com/SsrCoder/leetwatcher/output/output .

RUN upx -9 -o app output
####################################################
# 3. 运行
FROM alpine

# 镜像元数据，可以使用 docker inspect 查看
LABEL author="SsrCoder@gmail.com"

# EXPOSE 80

ENV APP_RUN_DIR /data

WORKDIR $APP_RUN_DIR

RUN apk update \
    && apk --no-cache add wget ca-certificates \
    && apk add -f --no-cache git \
    && apk add -U tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime

COPY --from=compress /data/app .
COPY conf ./conf

CMD ["./app"]