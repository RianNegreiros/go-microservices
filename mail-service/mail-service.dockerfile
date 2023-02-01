FROM alpine:latest

RUN mkdir /app

COPY mailer-service /app
COPY templates /templates

CMD [ "/app/mailer-service"]