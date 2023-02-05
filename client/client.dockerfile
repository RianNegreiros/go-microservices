FROM alpine:latest

RUN mkdir /app

COPY front-end /app

CMD [ "/app/front-end"]