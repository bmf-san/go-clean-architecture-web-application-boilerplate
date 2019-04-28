FROM golang:1.11.9-alpine3.9

WORKDIR /usr/local/project

RUN apk add --no-cache git
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/oxequa/realize

CMD realize start