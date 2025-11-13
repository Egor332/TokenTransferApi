# --- STAGE 1: Build ---
FROM golang:1.25.4-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git tzdata docker-cli

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main -ldflags="-w -s" ./cmd

# --- STAGE 2: Final Image ---
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache tzdata

COPY --from=builder /app/main .


EXPOSE 8080

CMD ["/app/main"]