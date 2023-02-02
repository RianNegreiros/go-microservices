FROM alpine:latest

RUN mkdir /app

COPY listener-service /app

CMD [ "/app/listener-service"]