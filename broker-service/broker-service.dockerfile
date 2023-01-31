FROM alpine:latest

RUN mkdir /app

COPY broker-service /app

CMD ["/app/broker-service"]