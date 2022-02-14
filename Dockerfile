FROM golang:1.17-alpine
WORKDIR /go/src/app

VOLUME /usr/src/myapp
VOLUME /go/pkg/mod
EXPOSE 8080
EXPOSE 8081
