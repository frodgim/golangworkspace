# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY . ./

RUN go mod tidy
RUN go build -o app

# Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
