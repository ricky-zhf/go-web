# 基础镜像
FROM golang:alpine AS builder
# 维护者
MAINTAINER ricky_chow
# 为镜像设置环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn
# docker build 时执行命令 - 创建目录
RUN mkdir -p "/app"
#&& ln -sf /dev/stdout /app/service.log \

# 工作目录
WORKDIR "/app"
# 将代码复制到容器中
COPY go.mod .
COPY go.sum .
# 安装依赖包
RUN go mod tidy
#RUN go mod download
# 拷贝
COPY . .
# 将我们的代码编译成二进制可执行文件app
RUN go env && go build -o user_server main.go

RUN ls -l


FROM debian:stretch-slim
# 需要再容器中使用sh命令，所以不能用stretch
# 从builder镜像中把go_app拷贝到当前目录下
COPY --from=builder /app/user_server /
COPY --from=builder /app/config ./config
# 拷贝config文件，目的文件目录没有config文件夹，所以要用./config
# 声明服务端口
EXPOSE 9091
# docker run 时执行命令
ENTRYPOINT ["/user_server"]
#docker build . -t user_server