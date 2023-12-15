# syntax=docker/dockerfile:1

FROM golang:1.21.5-alpine AS build

WORKDIR /app

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

COPY . .
COPY .env .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o /build

EXPOSE 8080

CMD [ "/build" ]