FROM alpine:3.9

ADD . /go/bin
WORKDIR /go/bin

CMD /go/bin/app

EXPOSE 9092