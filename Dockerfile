FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o finance_manager_api main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/finance_manager_api .
CMD ["./finance_manager_api"]

