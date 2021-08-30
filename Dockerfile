FROM golang:1.15.6-alpine3.12

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

RUN go build -o simple-go-service ./cmd/simple-go-service

CMD ["/app/simple-go-service"]
