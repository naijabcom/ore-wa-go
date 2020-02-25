FROM golang:alpine
LABEL maintainer="Nattapon Pondongnok <nainatjab999@gmail.com>"

WORKDIR /go/src/app

COPY ./src /go/src/app

RUN go mod download

CMD ["go", "run", "main.go"]
