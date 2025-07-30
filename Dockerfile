# Stage 1: Build
FROM golang:1.23 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the source code
COPY . .

# Build the Go app (binary will be named 'app')
RUN go build -o app

# ------------------------

# Stage 2: Run
FROM alpine:latest

# ทำให้ container crash ถ้ามี error
RUN apk --no-cache add ca-certificates

# Set working directory in runtime container
WORKDIR /root/

# Copy binary from builder stage
COPY --from=builder /app/app .

# บอกว่าต้องเปิดพอร์ตนี้
EXPOSE 8080

# Command to run
CMD ["./app"]