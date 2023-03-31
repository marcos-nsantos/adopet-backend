FROM golang:1.20.2-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o app ./cmd/server/main.go

FROM alpine:latest
COPY --from=builder /app .
CMD ["./app"]