FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git && apk add --update git mercurial && rm -rf /var/cache/apk/*
COPY . /go/src/github.com/eric1313/kcd
WORKDIR /go/src/github.com/eric1313/kcd

RUN export GO111MODULE=on && GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/kcd

#Actual Image
FROM alpine

VOLUME /go/src

RUN mkdir -p /kcd
ADD ./k8s /kcd/k8s/

COPY --from=builder /go/bin/kcd /kcd/
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*


WORKDIR /kcd
EXPOSE 2019
USER 1001

ENTRYPOINT ["/kcd/kcd"]
