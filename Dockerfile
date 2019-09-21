FROM alpine:latest

RUN apk update && apk add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY geo-tracking-redis .

CMD ./geo-tracking-redis