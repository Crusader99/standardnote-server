VERSION 0.7

build-all:
    BUILD --platform=linux/amd64 +build
    BUILD --platform=linux/arm64/v8 +build

build:
    ARG TARGETARCH
    ARG TARGETVARIANT
    FROM alpine
    ENV DATABASE_PATH /data/database
    RUN mkdir -p ${DATABASE_PATH}
    COPY (+compile-binary/bin --GOARCH=$TARGETARCH) /usr/local/bin/standardnote
    EXPOSE 5000
    CMD ["standardnote", "server", "-c", "/etc/standardfile/standardfile.yml"]
    SAVE IMAGE --push docker.io/crusaders/standardnote-server:latest

compile-binary:
    ARG GOARCH
    FROM --platform=linux/amd64 golang:1.20-alpine
    COPY --dir . /build
    ENV CGO_ENABLED 0
    ENV GO111MODULE on
    ENV GOPROXY https://proxy.golang.org
    ENV GOOS=linux
    WORKDIR /build/cmd/standardfile
    RUN go mod download
    RUN go build -o /build/bin
    SAVE ARTIFACT /build/bin

