FROM golang:1.22.4 AS builder

WORKDIR /app


COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /root/

COPY /config /config

COPY --from=builder /app/main .
EXPOSE 5001

CMD ["./main"]
