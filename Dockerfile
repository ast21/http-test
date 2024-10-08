FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY go.mod main.go ./
RUN go build -o main .

FROM scratch
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8000

CMD ["./main"]
