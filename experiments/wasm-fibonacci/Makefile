.DEFAULT_GOAL=run

wasm-build:
	GOOS=js GOARCH=wasm go build -o ./assets/wasm/fibonacci.wasm ./cmd/wasm-fibonacci

wasm-init:
# If it doesn't work, run directly in terminal
	cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./assets/wasm

server-run:
	go run ./cmd/server/main.go --dir ./assets

test:
	go test ./cmd/fibonacci

run: wasm-build server-run
