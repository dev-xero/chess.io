FROM golang:1.25.4-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server cmd/server/main.go

FROM alpine:latest
RUN adduser -D -h /app chess
WORKDIR /app
COPY --from=builder /app/server .
USER chess
EXPOSE 8080
CMD ["./server"]
