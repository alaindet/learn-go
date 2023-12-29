const CALC = {
  el: 'calc-start',
  path: 'wasm/calc.wasm',
  module: undefined,
  instance: undefined,
};

const $go = new Go(); // Comes from wasm_exec.js

// Load WASM for Calc
// Functions: add, subtract
WebAssembly.instantiateStreaming(fetch(CALC.path), $go.importObject).then(
  result => {
    CALC.module = result.module;
    CALC.instance = result.instance;
    (async () => await $go.run(CALC.instance))();
  }
);

function onAdd() {
  add('input-a', 'input-b', 'input-result');
}

function onSubtract() {
  subtract('input-a', 'input-b', 'input-result');
}
