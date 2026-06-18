# Build stage: compile the binary from source. The build context is the
# project root (see docker-compose.yml) so the shared "common" module is
# available here too, ready for when authentication starts importing it.
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY common ./common
COPY authentication ./authentication

WORKDIR /app/authentication
RUN CGO_ENABLED=0 go build -o authentication_app ./cmd/api

# Final stage: minimal runtime image with just the compiled binary.
FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/authentication/authentication_app /app
CMD ["/app/authentication_app"]
