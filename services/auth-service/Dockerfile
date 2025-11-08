# ----------------------------
# Build stage
# ----------------------------
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install git for go modules
RUN apk add --no-cache git

# Copy go.mod and go.sum and download dependencies
COPY go.* ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o auth-service ./cmd/main.go

# ----------------------------
# Runtime stage
# ----------------------------
FROM alpine:3.18

WORKDIR /app

# Install only ca-certificates
RUN apk add --no-cache ca-certificates

# Copy the binary from builder
COPY --from=builder /app/auth-service .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./auth-service"]
