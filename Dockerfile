ARG GO_VERSION=1.25.5

# =========================================================
# BUILD STAGE
# =========================================================

FROM registry.docker.com/library/golang:$GO_VERSION-alpine AS builder

WORKDIR /app

# Install packages needed to build
RUN apk add --no-cache git

# Download dependencies lebih cache-friendly
COPY go.mod go.sum ./

RUN go mod download

# Copy source
COPY . .

# Build application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" \
    -o /app/audit-trail-service \
    ./cmd/api

# =========================================================
# FINAL STAGE
# =========================================================

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/audit-trail-service .

COPY --from=builder /app/docs/ ./docs

ENTRYPOINT ["/app/audit-trail-service"]
