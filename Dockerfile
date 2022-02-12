FROM golang:1.17-alpine 
ENV GO111MODULE=on
WORKDIR ${GOPATH}/src/github.com/container-examples/golang-webserver/