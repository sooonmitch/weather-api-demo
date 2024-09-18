# Start with the Go Alpine image to build the binary
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./

FROM alpine:3.20  

ENV SERVER_ADDRESS=localhost:8080

COPY --from=builder /app/main /app/main

EXPOSE 8080

CMD ["/app/main"]