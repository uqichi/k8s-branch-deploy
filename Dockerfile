FROM golang:1.12.6-alpine3.10

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN apk add --no-cache alpine-sdk git

CMD ["go", "run", "main.go"]