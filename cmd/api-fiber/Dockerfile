# Dockerfile for Fiber API server
FROM golang:1.25 AS builder
WORKDIR /app
COPY ../../.. .
WORKDIR /app/cmd/api-fiber
RUN go build -o /app/bin/api-fiber main.go

FROM debian:trixie-slim
WORKDIR /app
COPY --from=builder /app/bin/api-fiber ./
EXPOSE 8080
CMD ["./api-fiber"]
