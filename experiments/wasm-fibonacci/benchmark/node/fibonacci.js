module.exports = function fibonacci(n) {
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
