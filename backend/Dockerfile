# Stage 1: build the Go binary
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Download dependencies first (leverages Docker layer cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy source and compile
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ticketapp .

# Stage 2: minimal runtime image
FROM alpine:3.19

WORKDIR /app

# Copy compiled binary from builder
COPY --from=builder /app/ticketapp .

EXPOSE 8080

ENTRYPOINT ["./ticketapp"]
