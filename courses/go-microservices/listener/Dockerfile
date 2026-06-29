FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY common ./common
COPY listener ./listener

WORKDIR /app/listener
RUN CGO_ENABLED=0 go build -o listener_app .

# Final stage: minimal runtime image with just the compiled binary.
FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/listener/listener_app /app
CMD ["/app/listener_app"]
