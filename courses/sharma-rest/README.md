## How to create a self-signed certificate for development
```
# Put "localhost" on Common Name
openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365
```

## How to run a benchmark
```
TODO
```

## Dependencies
- [wg/wrk](https://github.com/wg/wrk) HTTP benchmarking tool
- [rakyll/hey](https://github.com/rakyll/hey) HTTP benchmarking tool
- htop (optional)s
