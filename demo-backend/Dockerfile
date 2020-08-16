#FROM golang:1.14 AS builder
#
#ENV WORK_DIR=$GOPATH/src/github.com/sillyhatxu/remind-backend
#WORKDIR $WORK_DIR
#COPY . .
#RUN go mod download
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go
#
FROM xushikuan/alpine-build:1.0

WORKDIR /app

ENV TIME_ZONE=Asia/Singapore
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

RUN mkdir -p logs
COPY main .
ENTRYPOINT ./main