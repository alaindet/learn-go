# FROM golang:1.26-alpine AS builder
# RUN mkdir /app
# COPY . /app
# WORKDIR /app
# RUN CGO_ENABLED=0 go build -o broker_app ./cmd/api
# RUN chmod +x /app/broker_app

# FROM alpine:latest
# RUN mkdir /app
# COPY --from=builder /app/broker_app /app
# CMD ["/app/broker_app"]

# FROM alpine:latest
# FROM scratch
# FROM gcr.io/distroless/static
# COPY --from=builder /app/broker_app /broker_app
# ENTRYPOINT ["/broker_app"]

FROM alpine:latest
RUN mkdir /app
COPY broker_app /app
CMD ["/app/broker_app"]