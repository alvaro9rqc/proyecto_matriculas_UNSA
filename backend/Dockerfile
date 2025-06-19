FROM golang:latest AS builder

WORKDIR /app

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

RUN go mod download

COPY . /app

RUN go build -o main /app/cmd/main.go

FROM debian:bookworm-slim  AS production

WORKDIR /app

COPY --from=builder /app/db/database.db /app/db/database.db
COPY --from=builder /app/main /app/main

CMD ["/app/main"]
