# --- Etapa 1: build ---
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ticketapp .

# --- Etapa 2: imagen final ---
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/ticketapp .

EXPOSE 8080

CMD ["./ticketapp"]
