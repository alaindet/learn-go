# Let's Go Further companion repo

This repo is used to learn how to create a solid JSON RESTful API in Go via Alex Edward's book "Let's Go Further".

## How to run

### Code in development mode
```
go run ./cmd/api
```

### Database
To start the database
```
docker run --name greenlight-db \
  -e POSTGRES_USER=admin \
  -e POSTGRES_PASSWORD=admin \
  -e POSTGRES_DB=greenlight \
  -p 5432:5432 \
  -v greenlight_pgdata:/var/lib/postgresql/data \
  -d postgres:16-alpine
```

To run `psql`
```
docker exec -it greenlight-db psql -U admin -d greenlight
```