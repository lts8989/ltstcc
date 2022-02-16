FROM golang:alpine3.15
RUN apk add protobuf && \
    go env -w GO111MODULE=on  && \
      go env -w GOPROXY=https://goproxy.cn,direct && \
      go env && \
      go get google.golang.org/protobuf/cmd/protoc-gen-go && \
      go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

