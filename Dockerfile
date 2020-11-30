FROM golang:alpine

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOPRIVATE=github.com/schigh-ntwrk/*

RUN apk update && apk upgrade && apk add --no-cache make git bash
RUN rm -rf /var/cache/apk/* /tmp/*
