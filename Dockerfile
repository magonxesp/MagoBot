FROM golang:1.19-alpine

WORKDIR /app

RUN apk add --no-cache make

COPY . .

RUN make build

CMD ["/app/build/magobot"]