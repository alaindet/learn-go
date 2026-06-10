FROM golang:1.26-alpine AS builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api
RUN chmod +x /app/brokerApp

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/brokerApp /app
CMD ["/app/brokerApp"]

# FROM alpine:latest
# FROM scratch
# FROM gcr.io/distroless/static
# COPY --from=builder /app/brokerApp /brokerApp
# ENTRYPOINT ["/brokerApp"]