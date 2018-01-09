# Builder
FROM golang:latest as builder

COPY . /go/src/github.com/damianopetrungaro/golang-bookshop
WORKDIR /go/src/github.com/damianopetrungaro/golang-bookshop

RUN curl https://glide.sh/get | sh && \
    go get github.com/rubenv/sql-migrate/... && \
    glide install && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/bookshop .

# App
FROM alpine:latest

ENV SQL_CONNECTION_URL postgres://root:root@postgresql/bookshop?sslmode=disable

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/bin/bookshop bookshop
ENTRYPOINT ./bookshop