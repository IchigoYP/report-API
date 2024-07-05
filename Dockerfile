FROM golang:1.20.4 as build

CMD ["/bin/bash", "-b"]

ENV GOPATH=/go

RUN apt-get update && apt-get upgrade -y\
    && apt-get autoremove -y\
    && apt-get clean\
    && rm -rf /var/lib/apt/lists/*

WORKDIR /go/src

COPY ./src/go.mod /go/src/go.mod
COPY ./src/go.sum /go/src/go.sum
RUN go mod download
