FROM golang:alpine3.15
RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.7/main/ > /etc/apk/repositories
RUN apk update \
        && apk add tzdata protobuf \
        && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
        && echo "Asia/Shanghai" > /etc/timezone \
      && go env -w GO111MODULE=on  && \
      go env -w GOPROXY=https://goproxy.cn,direct && \
      go env && \
      go get google.golang.org/protobuf/cmd/protoc-gen-go && \
      go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

