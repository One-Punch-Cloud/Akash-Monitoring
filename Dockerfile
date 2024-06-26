FROM golang:1.15 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o onePunchAkashMonitoring ./cmd/main

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/onePunchAkashMonitoring .
CMD ["./onePunchAkashMonitoring", "/onePunchAkashMonitoring/config.json"]