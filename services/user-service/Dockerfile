# ----------------------------
# Build stage
# ----------------------------
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install git for go modules
RUN apk add --no-cache git

# Copy go.mod and go.sum and download dependencies
COPY services/user-service/go.* ./
COPY cosmix-events /cosmix-events
RUN go mod download

# Copy the rest of the code
COPY services/user-service/ .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o user-service ./cmd/main.go

# ----------------------------
# Runtime stage
# ----------------------------
FROM alpine:3.18

WORKDIR /app

# Install only ca-certificates
RUN apk add --no-cache ca-certificates

# Copy the binary from builder
COPY --from=builder /app/user-service .

# Expose port
EXPOSE 8081

# Run the binary
CMD ["./user-service"]
