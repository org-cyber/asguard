# ---- Builder stage ----
FROM golang:1.25-bookworm AS builder

WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o general-engine ./cmd/server/

# ---- Runtime stage ----
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/general-engine /app/general-engine

EXPOSE 8083

CMD ["/app/general-engine"]