.DEFAULT_GOAL=run

clear:
	@rm -rf generated
	@rm -f main

generate: clear
	@go run cmd/generate/generate.go

build: generate
	@go build cmd/main/main.go

run: build
	@./main
