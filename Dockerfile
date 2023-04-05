FROM golang:1.20.3-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o app ./cmd/server/main.go

FROM scratch
COPY --from=builder /app .
CMD ["./app"]