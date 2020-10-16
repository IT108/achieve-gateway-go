# syntax = docker/dockerfile:1-experimental

FROM golang:1.15.2-buster AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=1

WORKDIR /go/src/github.com/achieve-gateway-go
COPY . .
RUN go mod init
RUN go mod tidy

RUN go get -d -v ./...
RUN go install -v ./...

ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=1 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /app/main .

WORKDIR /app
RUN ldd main | tr -s '[:blank:]' '\n' | grep '^/' | \
    xargs -I % sh -c 'mkdir -p $(dirname ./%); cp % ./%;'
RUN mkdir -p lib64 && cp /lib64/ld-linux-x86-64.so.2 lib64/


FROM scratch

COPY --chown=0:0 --from=builder /app /

ENTRYPOINT ["/main"]