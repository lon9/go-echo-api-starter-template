FROM golang:alpine

RUN apk add --no-cache git gcc musl-dev
ADD . /go/src/github.com/lon9/go-echo-api-starter
WORKDIR /go/src/github.com/lon9/go-echo-api-starter

RUN go get
