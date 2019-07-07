FROM golang:1.11.9-alpine3.9

WORKDIR /go/src/github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app

RUN apk add --no-cache git \
    binutils-gold \
    curl \
    g++ \
    gcc \
    gnupg \
    libgcc \
    linux-headers \
    make
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/oxequa/realize
RUN go get -u github.com/go-chi/chi

CMD realize start