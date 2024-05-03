FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o sirs-app ./cmd/api

# Path: config/Dockerfile.backend

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/sirs-app /app/sirs-app

ENTRYPOINT [ "/app/sirs-app" ] 