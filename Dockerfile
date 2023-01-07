FROM golang:1.16-alpine

WORKDIR /app

ADD go.mod ./

RUN go mod download

ADD client ./
