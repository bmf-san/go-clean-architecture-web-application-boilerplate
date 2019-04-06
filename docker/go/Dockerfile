FROM golang:1.10.7-alpine3.8

WORKDIR /usr/local/project

RUN apk add --no-cache git
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/oxequa/realize

CMD realize start