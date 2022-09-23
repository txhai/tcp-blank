FROM golang:1.18-alpine3.15 as builder
WORKDIR /app
ADD . .
RUN go build -o /app/tcpblank main.go

FROM alpine:3.15
COPY --from=builder /app/tcpblank /usr/bin
ENTRYPOINT [ "/usr/bin/tcpblank" ]