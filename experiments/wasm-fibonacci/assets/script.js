const FIBONACCI = {
  path: 'wasm/fibonacci.wasm',
  module: undefined,
  instance: undefined,
  threshold: 93,
};

const $go = new Go(); // Comes from wasm_exec.js

// Load WASM for Fibonacci
// Functions: fibonacci
WebAssembly.instantiateStreaming(fetch(FIBONACCI.path), $go.importObject).then(
  result => {
    FIBONACCI.module = result.module;
    FIBONACCI.instance = result.instance;
    (async () => await $go.run(FIBONACCI.instance))();
  }
);

function onFibonacci() {
  const input = document.getElementById('input');
  const output = document.getElementById('output');

  const n = parseInt(input.value);

  if (n > FIBONACCI.threshold) {
    console.log(`N must be smaller than ${FIBONACCI.threshold}`);
    return;
  }

  output.value = fibonacci(n);
}

function onFibonacciBenchmark() {
  console.log('');
  measureTime("JavaScript", () => {
    for (let i = 0; i < 1000000; i++) {
      jsFibonacci(75);
    }
  });
  measureTime("WASM", () => {
    for (let i = 0; i < 1000000; i++) {
      fibonacci(75);
    }
  });
}

function measureTime(name, fn) {
  const start = performance.now();
  fn();
  const took = performance.now() - start;
  console.log(`Measure time: "${name}" took ${took} ms`);
}

function jsFibonacci(n) {
	if (n == 0 || n == 1) {
		return n;
	}

	let secondLast = 0;
	let last = 1;
	let result = 1;

	for (let i = 2; i <= n; i++) {
		result = secondLast + last
		secondLast = last
		last = result
	}

	return result;
}
