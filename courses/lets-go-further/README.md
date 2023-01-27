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
