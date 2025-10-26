FROM golang:1.25.3

WORKDIR /gobook

COPY go.mod ./

COPY . .
