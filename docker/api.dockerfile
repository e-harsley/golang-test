FROM golang:alpine

WORKDIR /go-test/cmd
ADD . /go-test/cmd

RUN go mod download

ENTRYPOINT go build  && ./cmd