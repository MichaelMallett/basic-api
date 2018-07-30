FROM golang:1.9-alpine
RUN apk add --no-cache ca-certificates \
        dpkg \
        gcc \
        git \
        musl-dev \
        bash

RUN go get github.com/tockins/realize