# =========================
# Stage 1: Builder
# =========================
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Set build env
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy dependency dulu untuk cache
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -ldflags="-s -w" -o app


# =========================
# Stage 2: Runtime
# =========================
FROM alpine:3.19

WORKDIR /app

# Copy binary saja
COPY --from=builder /app/app .

# Default env (bisa dioverride saat run)
ENV APP_ENV=production \
    PORT=3000

EXPOSE 3000

CMD ["./app"]
