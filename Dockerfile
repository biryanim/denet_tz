FROM golang:1.24-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git make bash
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build install-deps

FROM alpine:3.21
WORKDIR /app

RUN apk add --no-cache bash make

COPY --from=builder /app/bin /app/bin
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /app/local.env /app/local.env
COPY --from=builder /app/Makefile /app/Makefile

EXPOSE 8080

CMD sh -c "make migration-up && /app/bin/main"