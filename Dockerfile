# Builder
FROM golang:latest as builder

COPY . /go/src/github.com/damianopetrungaro/golang-bookshop
WORKDIR /go/src/github.com/damianopetrungaro/golang-bookshop

RUN curl https://glide.sh/get | sh && \
    glide install && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/bookshop .

# App
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/bin/bookshop bookshop
ENTRYPOINT ./bookshop