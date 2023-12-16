# syntax=docker/dockerfile:1

FROM golang:1.21.5-alpine AS build

RUN mkdir /app
WORKDIR /app

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

COPY . .
COPY .env .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon


ENTRYPOINT CompileDaemon -build="go build -o api" -command="./api"
