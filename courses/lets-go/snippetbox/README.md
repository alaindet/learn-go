# Snippetbox

## Usage

```
cd ./snippetbox
cp .env.example .env # Then fill the file
go mod download # Equivalent to `npm install`
go mod verify
cd ./tls
go run generate_cert.go --rsa-bits=2048 --host=localhost
cd ..
go run ./cmd/web
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

## Testing

For integration and e2e tests, run test containers via Docker
```
docker-compose -f ./docker-compose.test.yaml up -d
docker-compose down
```

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

## Skip long tests
Tests are marked as "long" or "short" by the user via `testing.Short()` and `testing.T.Skip()`
```
go test -short ./...
```

## Check code coverage
```
go test -cover ./...

# Export a detailed report
go test -coverprofile=./coverage-profile.dat ./...

# See report in terminal
go tool cover -func=./coverage-profile.dat

# See report as HTML in browser
go tool cover -html=./coverage-profile.dat
```
