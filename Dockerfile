FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY cmd ./cmd
RUN go build -o /bin/server ./cmd/server

FROM alpine:3.19
COPY --from=builder /bin/server /server
EXPOSE 8080
CMD ["/server"]
