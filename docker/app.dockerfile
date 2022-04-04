FROM golang:alpine

WORKDIR /novocaine-dev

ADD . .

RUN go mod download

ENTRYPOINT go build  && ./novocaine-dev