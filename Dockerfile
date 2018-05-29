FROM golang:1.9

ADD . /go/src/github.com/OdaDaisuke/emo-lyrics-api

WORKDIR /go/src/github.com/OdaDaisuke/emo-lyrics-api↲

RUN go get -u github.com/golang/dep/cmd/dep↲
RUN dep ensure
RUN go install

EXPOSE 80↲
