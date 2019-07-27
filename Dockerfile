FROM golang:1.12.4-alpine
MAINTAINER mnhan0403@gmail.com

RUN apk update && apk upgrade && \
  apk add --no-cache bash git openssh

WORKDIR $GOPATH/src/github.com/smu-gp/sp-auth-server
COPY . .

ENV GO111MODULE on
RUN go get -u
RUN go build -o main main.go

EXPOSE 8001

CMD ["./main"]