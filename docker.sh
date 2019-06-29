#!/bin/bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo  -ldflags '-w -s' -o app .
upx app
docker build -t go-hosts .
docker run -it -p 9092:9092 go-hosts
