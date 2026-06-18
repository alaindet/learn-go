# Build stage: compile the binary from source so the
# "replace common => ../common" directive in broker/go.mod can resolve.
# The build context is the project root (see docker-compose.yml).
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy the shared module first; broker/go.mod's "replace common => ../common"
# resolves to /app/common once we build from /app/broker.
COPY common ./common
COPY broker ./broker

WORKDIR /app/broker
RUN CGO_ENABLED=0 go build -o broker_app ./cmd/api

# Final stage: minimal runtime image with just the compiled binary.
FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/broker/broker_app /app
CMD ["/app/broker_app"]
