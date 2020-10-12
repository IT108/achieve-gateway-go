# syntax = docker/dockerfile:1-experimental

FROM golang:1.15.2-buster AS builder

WORKDIR /go/src/github.com/achieve-gateway-go
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /app/main .


ENTRYPOINT ["/app/main"]