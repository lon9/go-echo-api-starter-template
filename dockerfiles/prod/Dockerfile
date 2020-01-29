FROM golang:alpine as build-env

ARG GITHUB_ACCESS_TOKEN

RUN apk add --no-cache git gcc musl-dev
RUN git config --global url."https://${GITHUB_ACCESS_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

RUN go get github.com/lon9/go-echo-api-starter
WORKDIR /go/src/github.com/lon9/go-echo-api-starter
RUN go build -o /usr/bin/app

FROM alpine
COPY --from=build-env /usr/bin/app /usr/bin/app

