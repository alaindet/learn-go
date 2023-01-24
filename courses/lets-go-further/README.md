# Greenlight

## Usage

```
cd ./greenlight
cp .env.example .env # Then fill the file
go mod download # Equivalent to `npm install`
go mod verify
# TODO
```

## Start development
```
cd ./greenlight && \
export $(grep -v '^#' .env | xargs -d '\n') && \
go run ./cmd/api
```

## Check the available CLI flags
```
go run ./cmd/api -help
```

## Generate TLS key pair
```
```

## Build
```
```

## Testing
```
```
