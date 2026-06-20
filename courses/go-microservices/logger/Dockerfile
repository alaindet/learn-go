FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY common ./common
COPY logger ./logger

WORKDIR /app/logger
RUN CGO_ENABLED=0 go build -o logger_app ./cmd/api

# Final stage: minimal runtime image with just the compiled binary.
FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/logger/logger_app /app
CMD ["/app/logger_app"]
