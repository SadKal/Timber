FROM golang:1.22.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/ .
EXPOSE 8080
CMD ["./main"]