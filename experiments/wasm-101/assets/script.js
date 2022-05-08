/**
From wasm_exec.js
=================
Go class
WebAssembly class
*/

const goWasm = new Go();

WebAssembly.instantiateStreaming(fetch("wasm/main.wasm"), goWasm.importObject)
  .then(result => {
    goWasm.run(result.instance);
  });
