FROM golang:alpine

LABEL maintainer="Gaby"

WORKDIR /app

COPY . .

RUN go build -o dockerize

RUN apk update && apk add bash

EXPOSE 8080

CMD ["./dockerize"]
