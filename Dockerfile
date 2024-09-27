FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o chat-app ./cmd/server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/chat-app .
COPY .env .
COPY web /app/web
EXPOSE 80
CMD ["./chat-app"]
