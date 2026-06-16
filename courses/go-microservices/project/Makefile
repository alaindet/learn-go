AUTHENTICATION_BINARY=authentication_app
BROKER_BINARY=broker_app
FRONTEND_BINARY=frontend_app

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_authentication
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_authentication: builds the authentication binary as a linux executable
build_authentication:
	@echo "Building authentication binary..."
	cd ../authentication && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTHENTICATION_BINARY} ./cmd/api
	@echo "Done. Built authentication binary"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ../broker && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done. Built broker binary"

## build_frontend: builds the froneend binary
build_frontend:
	@echo "Building frontend binary..."
	cd ../frontend && env CGO_ENABLED=0 go build -o ${FRONTEND_BINARY} ./cmd/web
	@echo "Done. Built frontend binary"

## start: starts the frontend
start: build_frontend
	@echo "Starting frontend"
	cd ../frontend && ./${FRONTEND_BINARY} &

## stop: stop the frontend
stop:
	@echo "Stopping frontend..."
	@-pkill -SIGTERM -f "./${FRONTEND_BINARY}"
	@echo "Stopped frontend!"