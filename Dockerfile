# Builder
FROM golang:1.22-alpine3.19 AS builder
ENV TZ=Asia/Jakarta
RUN mkdir /app
ADD . /app

WORKDIR /app

COPY  go.mod  .

RUN go mod tidy

RUN go build -o engine ./app/

EXPOSE 9090

CMD [ "/app/engine", "rest" ]