FROM golang:alpine AS build_base

LABEL maintainer="Nattapon Pondongnok <nainatjab999@gmail.com>"

ENV GO111MODULE=on

WORKDIR /go/src/app

COPY . /go/src/app

RUN go mod download

CMD ["go", "run", "main.go"]
