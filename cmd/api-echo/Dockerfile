# Dockerfile for Echo API server
FROM golang:1.25 AS builder
WORKDIR /app
COPY ../../.. .
WORKDIR /app/cmd/api-echo
RUN go build -o /app/bin/api-echo main.go

FROM debian:trixie-slim
WORKDIR /app
COPY --from=builder /app/bin/api-echo ./
EXPOSE 8080
CMD ["./api-echo"]
