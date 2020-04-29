FROM golang:1.14-alpine

RUN apk add build-base
RUN apk --no-cache add apache2-utils