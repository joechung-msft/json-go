# Dockerfile for Gin API server
FROM golang:1.25 AS builder
WORKDIR /app
COPY ../../.. .
WORKDIR /app/cmd/api-gin
RUN go build -o /app/bin/api-gin main.go

FROM debian:trixie-slim
WORKDIR /app
COPY --from=builder /app/bin/api-gin ./
EXPOSE 8080
CMD ["./api-gin"]
