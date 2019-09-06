FROM golang:1.12.7-alpine

WORKDIR $GOPATH/src/simple-go-rest-service

ADD ../.. $GOPATH/src/simple-go-rest-service

RUN apk add git
RUN apk add build-base

RUN go get -d -v ./...

RUN go get github.com/stretchr/testify/assert
RUN go get github.com/golangci/golangci-lint/cmd/golangci-lint

RUN go test

RUN go install -v ./...

RUN golangci-lint run