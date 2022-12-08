# Snippetbox

## Usage

```
cd ./snippetbox &&
go run ./cmd/web # Use default in .env

go run ./cmd/web -addr=":9999"
go run ./cmd/web -addr ":9999" # Equivalent
export $(cat .env | xargs) && go run ./cmd/web -addr=$APP_ADDRESS
go run ./cmd/web -help
```
