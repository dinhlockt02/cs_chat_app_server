FROM golang:1.19-bullseye as builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main

FROM debian:bullseye-slim

RUN apt-get update

RUN apt-get install -y ca-certificates

RUN update-ca-certificates

WORKDIR /app

COPY --from=builder /app/main .
COPY ./templates /app/templates/

ENV SERVER_PORT=8080
ENV GIN_MODE=release

EXPOSE 8080

CMD ["/app/main"]
