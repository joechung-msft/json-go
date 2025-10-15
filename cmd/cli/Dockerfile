# Dockerfile for CLI
FROM golang:1.25 AS builder
WORKDIR /app
COPY ../../.. .
WORKDIR /app/cmd/cli
RUN go build -o /app/bin/json-cli main.go

FROM debian:trixie-slim
WORKDIR /app
COPY --from=builder /app/bin/json-cli ./
CMD ["./json-cli"]
