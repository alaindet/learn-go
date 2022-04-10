BINARY_NAME=myapp

build:
	@go mod vendor
	@echo "Building Gomitolo..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Gomitolo built!"

run: build
	@echo "Starting Gomitolo..."
	@./tmp/${BINARY_NAME} &
	@echo "Gomitolo started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Gomitolo..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Gomitolo!"

restart: stop start
