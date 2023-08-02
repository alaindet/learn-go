# Greenlight

## Usage

```console
cd ./greenlight
cp .env.example .env
# Fill the new .env
export $(grep -v '^#' .env | xargs -d '\n') # Export all vars from .env
go mod download # Equivalent to `npm install`
go mod verify
docker-compose up -d
# Install migrate
migrate -path=./database/migrations -database=$GREENLIGHT_DB_DSN up
go run ./cmd/api # Run it in development
# Or build it (todo)
curl http://localhost:4000/v1.0/healthcheck | json_pp
```

## Start development
```console
cd ./greenlight && \
docker-compose up -d && \
export $(grep -v '^#' .env | xargs -d '\n') && \
go run ./cmd/api
```

## Stop development
```console
docker-compose down
```

## Check the available CLI flags
```console
go run ./cmd/api -help
```

## Generate TLS key pair
```console
```

## Build
```console
```

## Testing
```console
```

## 3rd-party binaries installed

- `golang-migrate`
  ```console
  curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
  mv migrate.linux-amd64 $GOPATH/bin/migrate
  ```

- `hey` (todo)
- `caddy` (todo)

## Migrations

- Create a migration (creates `*.up.sql` and `*.down.sql` coupled files)
  ```console
  migrate create -seq -ext=.sql -dir=./database/migrations create_movies_table
  ```
- Run all migrations
  ```console
  cd ./greenlight && \
  export $(grep -v '^#' .env | xargs -d '\n') && \
  migrate -path=./database/migrations -database=$GREENLIGHT_DB_DSN up
  ```

## Open Bash on running Docker container
```console
docker exec -it <container name> /bin/bash
```

## Execute command on running Docker container
```console
docker exec -it <container name> <command>
```

## Open Bash on database service
```console
docker-compose run db /bin/bash
```
