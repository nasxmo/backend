# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /ory-backend ./cmd/api

# Runtime stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /ory-backend /app/ory-backend
COPY --from=builder /app/pkg/config/config.yaml /app/config.yaml

EXPOSE 8080

CMD ["/app/ory-backend"]