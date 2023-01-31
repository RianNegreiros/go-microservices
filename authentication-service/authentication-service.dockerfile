FROM alpine:latest

RUN mkdir /app

COPY auth-service /app

CMD ["/app/auth-service"]