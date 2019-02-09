FROM golang:1.11

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/anuragdhingra/go-rest-example

ADD . /go/src/github.com/anuragdhingra/go-rest-example

RUN dep ensure