FROM golang:latest

COPY . /go/src/bookshop
WORKDIR /go/src/bookshop

RUN curl https://glide.sh/get | sh && \
    go get github.com/go-task/task/cmd/task

ENTRYPOINT task deps build

EXPOSE 80