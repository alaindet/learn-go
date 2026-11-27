# Let's Go Further companion repo

This repo is used to learn how to create a solid JSON RESTful API in Go via Alex Edward's book "Let's Go Further".

## Dependencies
- Dependencies are Go-based tools that are listed as "tool dependencies" inside the project
- To use them, run `go tool hey` or `go tool migrate`

## How to run

### Code in development mode
```
go run ./cmd/api -db-dsn="postgres://greenlight_usr:greenlight_pwd@localhost:5432/greenlight_db?sslmode=disable"
```

### Database

Control database container
```
# Start
docker compose up -d

# Stop
docker compose down

# Show logs
docker compose logs -f postgres

# Reset database data by wiping the Docker volume
docker compose down -v

# Run migrations
# This is because golang-migrate cannot currently run as a "tool dependency" due to build tags not being accepted
go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate \
  -path=./database/migrations \
  -database=postgres://greenlight_usr:greenlight_pwd@localhost:5432/greenlight_db?sslmode=disable \
  up
```

Feed dsn to app
```
# Directly
go run ./cmd/api -db-dsn="postgres://greenlight_usr:greenlight_pwd@localhost:5432/greenlight_db?sslmode=disable"

# View env
export GREENLIGHT_DB_DSN="postgres://greenlight_usr:greenlight_pwd@localhost:5432/greenlight_db?sslmode=disable"
go run ./cmd/api
```

How to execute `psql` commands
```
# This
psql -d greenlight_db

# Becomes this
docker exec -it greenlight_container psql -U greenlight_usr -d greenlight_db
```

Migrations
```
# Apply migrations UP
migrate -path=./migrations -database="postgres://greenlight_usr:greenlight_pwd@localhost:5432/greenlight_db?sslmode=disable" up

# Apply migrations DOWN
migrate -path=./migrations -database="postgres://greenlight_usr:greenlight_pwd@localhost:5432/greenlight_db?sslmode=disable" down
```
