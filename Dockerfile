FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go build -o heaven-api main.go

FROM debian:stable-slim

WORKDIR /app

COPY --from=builder /app/heaven-api .

EXPOSE 3450

CMD ["./heaven-api"]