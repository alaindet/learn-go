.DEFAULT_GOAL=run

wasm-build:
	GOOS=js GOARCH=wasm go build -o ./assets/wasm/main.wasm ./cmd/wasm

wasm-init:
# If it doesn't work, run directly in terminal
	cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./assets/wasm

server-run:
	node server.js

run: wasm-build server-run
