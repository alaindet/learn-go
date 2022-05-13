module.exports = function (name, fn) {
  const start = performance.now();
  fn();
  const took = (performance.now() - start).toFixed(0);
  console.log(`Measure time: "${name}" took ${took} ms`);
}
