# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

RUN apk add git

WORKDIR /app

ENV GO111MODULE=on
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
RUN go mod download

COPY ./ ./

RUN go build -o /docker-gs-ping

CMD [ "/docker-gs-ping" ]
