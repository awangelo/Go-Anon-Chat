FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o chat-app .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/chat-app .
EXPOSE 8080
CMD ["./chat-app"]
