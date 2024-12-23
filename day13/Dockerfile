FROM golang:latest AS builder

WORKDIR /app
COPY main.go .
RUN go mod init hello-world-go
RUN go build -o main .

# Final stage
FROM alpine:3.18

WORKDIR /app
COPY --from=builder /app/main .

ENV PORT=8080
ENV SECRET_MESSAGE="default secret"

EXPOSE 8080

CMD ["./main"]
