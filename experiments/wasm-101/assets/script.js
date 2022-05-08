const goWasm = new Go(); // Comes from wasm_exec.js

const run = (mainWasm) => {
  // Run the module
  goWasm.run(mainWasm.instance);

  // Use a global function
  const btn = document.getElementById("clickme");
  const out = document.getElementById("output");

  btn.addEventListener("click", () => {
    const fromWasm = sayHello();
    console.log('From WASM', typeof fromWasm, fromWasm);
    out.innerHTML = fromWasm;
  });
};

(async () => {
  const w = WebAssembly.instantiateStreaming;
  run(await w(fetch("wasm/main.wasm"), goWasm.importObject));
})();
