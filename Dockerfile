FROM golang:1.15 AS build

ADD ./ /go/src/URL_shortener
WORKDIR /go/src/URL_shortener

RUN go install
