FROM golang:1.23.3-alpine3.20 AS base
WORKDIR /app
COPY /app .

RUN go install github.com/codegangsta/gin@latest
RUN go mod download
RUN apk add moreutils

FROM base AS builder
WORKDIR /app
COPY /app .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun

FROM scratch
WORKDIR /app
COPY --from=builder /app/cloudrun .
ENTRYPOINT ["./cloudrun"]