# --- build stage ---
FROM golang:1.22-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git

# Cho phép Go 1.22 tự tải toolchain phù hợp (1.24.x) nếu go.mod yêu cầu
ENV GOTOOLCHAIN=auto

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy code và build
COPY . .
RUN go build -o server ./cmd/server

# --- runtime stage ---
FROM alpine:3.19
WORKDIR /root
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]