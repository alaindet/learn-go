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
