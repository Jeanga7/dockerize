FROM golang:1.20.5

LABEL mainteners="jean"
LABEL version="1.0"

WORKDIR /app

COPY . .

RUN go build -o ascii

EXPOSE 8080

CMD ["./ascii"]
