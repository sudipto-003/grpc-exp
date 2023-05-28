FROM golang:1.20-alpine AS builder
WORKDIR /go/src/unary
ENV GO111MODULE=on

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

COPY ./proto/ ./proto/
COPY ./server/ ./server/

RUN go build -o /go/src/unary/bin/server ./server

FROM alpine:3.18
WORKDIR /app

COPY --from=builder /go/src/unary/bin/server ./

EXPOSE 5000
ENTRYPOINT ["./server"]