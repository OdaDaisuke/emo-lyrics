FROM golang:1.9

RUN go get -u gopkg.in/godo.v2/cmd/godo

WORKDIR /go/src/app

CMD ["/go/bin/godo", "server", "--watch"]

EXPOSE 8080