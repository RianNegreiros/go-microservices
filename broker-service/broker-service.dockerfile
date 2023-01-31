FROM golang:1.19-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o broker-service ./cmd/api

RUN chmod +x /app/broker-service

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/broker-service /app

CMD ["/app/broker-service"]