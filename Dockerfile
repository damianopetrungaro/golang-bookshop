FROM golang:latest

ENV ENV_NAME=production

ADD . /go/src/github.com/damianopetrungaro/golang-bookshop
WORKDIR /go/src/github.com/damianopetrungaro/golang-bookshop

RUN curl https://glide.sh/get | sh && \
    go get github.com/go-task/task/cmd/task && \
    task deps build

# If is production rm /go/src

ENTRYPOINT /go/bin/bookshop

EXPOSE 80