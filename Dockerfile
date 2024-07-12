FROM golang:1.25-alpine AS builder
WORKDIR /usr/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

FROM alpine
COPY --from=builder /usr/src/app/main ./main
CMD ["./main"]