const fibonacci = require('./fibonacci');
const measureTime = require('./measure-time');

const N = parseInt(process.env.N);
const ITERATIONS = parseInt(process.env.ITERATIONS);

measureTime('Node Fibonacci', () => {
  for (let i = 0; i < ITERATIONS; i++) {
    fibonacci(N);
  }
});
