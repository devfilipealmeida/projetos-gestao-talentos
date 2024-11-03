FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main .

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /main /app/main
COPY .env /app/.env
COPY migrations /app/migrations

RUN chmod +x /app/main

EXPOSE 8080

CMD ["./main"]