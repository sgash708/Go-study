FROM golang:1.15.7

RUN mkdir /go/src/go-study
WORKDIR /go/src/go-study
# Docker内でパッケージ管理させる記述
ENV GO111MODULE=on
RUN go mod download

ADD . /go/src/go-study
