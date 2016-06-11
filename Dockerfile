FROM golang:1.5

ENV GOBIN /go/bin
ENV GOPATH /go/src/github.com/omie/messages/Godeps/_workspace/:/go

ADD . /go/src/github.com/omie/messages
WORKDIR /go/src/github.com/omie/messages

RUN go install

CMD ["/go/bin/messages"]

