# Stage 1: Build
FROM golang:1.26 AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd

# Stage 2: Run
FROM alpine:latest

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/app .

# Expose port (change if needed)
EXPOSE 8080

CMD ["./app"]