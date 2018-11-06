## build stage
FROM golang:1.10-alpine AS builder
WORKDIR /go/src/github.com/rhasan33/goplate
COPY . .

RUN set -e
RUN apk add git
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -o dist/app


## certs stage, need openssl as well?
FROM alpine:latest as certs
RUN apk --update add ca-certificates


## final stage
FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/rhasan33/goplate/dist/app /app

LABEL Name=book_reader Version=0.0.1
EXPOSE 5000

ENTRYPOINT [ "/app" ]
CMD ["serve"]
