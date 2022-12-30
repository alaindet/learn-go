# Snippetbox

## Usage

```
cd ./snippetbox
cp .env.example .env
# Fill the file
go mod download # Equivalent to `npm install`
go mod verify
go run ./cmd/web

## One liner
cd ./snippetbox && go run ./cmd/web
```

## Generate TLS key pair
```
cd ./tls
go run generate_cert.go --rsa-bits=2048 --host=localhost
```

## Build
```
go build -o ./tmp/web ./cmd/web
cp -r ./tls ./tmp/
cp .env ./tmp/.env
cd ./tmp
./web
```

## Test

## All tests in the module
```
go test ./...
```

## All tests in `cmd/web`
```
go test ./cmd/web
```

## Only tests matching a regex
```
go test -run="^TestPing$" ./...
go test -run="^TestPing$" ./cmd/web
```

## Only subtests matching `{test regex}/{subtest regex}`
```
go test -run="^TestFriendlyDate$/^UTC$" ./...
go test -run="^TestFriendlyDate$/^UTC$" ./cmd/web
```

## Force all tests (no cached tests)
```
go test -count=1 ./...
go test -count=1 ./cmd/web

# Or
go clean -testcache
go test ./...
```

## Race detection
```
go test -race ./...
go test -race ./cmd/web
```
