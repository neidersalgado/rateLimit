FROM golang:1.18 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ratelimit ./cmd/ratelimit
FROM scratch

COPY --from=builder /app/ratelimit .

EXPOSE 8080

CMD ["./rateLimit"]
