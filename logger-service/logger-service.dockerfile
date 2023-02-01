FROM alpine:latest

RUN mkdir /app

COPY logger-service /app

CMD [ "/app/logger-service"]