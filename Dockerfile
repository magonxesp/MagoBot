FROM golang:1.24-alpine AS builder

WORKDIR /build

RUN apk add --no-cache make

COPY . .

RUN go build

FROM golang:1.24-alpine

WORKDIR /app

COPY --from=builder /build/MagoBot /app/MagoBot
RUN chmod +x /app/MagoBot

CMD ["/app/MagoBot"]