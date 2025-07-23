const { performance } = require('perf_hooks');

// Simple memoization cache
const fibCache = new Map();

function simulateWork(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function fibonacci(n) {
    if (n < 0) {
        throw new Error("Input must be a non-negative integer.");
    }
    if (n === 0) return 0;
    if (n === 1) return 1;

    if (fibCache.has(n)) {
        return fibCache.get(n);
    }

    await simulateWork(1); // Simulate work

    const result = fibonacci(n - 1) + fibonacci(n - 2);
    fibCache.set(n, result);
    return result;
}

async function calculateConcurrentFibonacci(maxN, numThreads) {
    console.log(`\n--- JavaScript (Node.js) Example ---`);
    console.log(`JavaScript: Calculating Fibonacci numbers up to ${maxN} concurrently using simulated async...`);
    const startTime = performance.now();

    const promises = [];
    for (let i = 0; i <= maxN; i++) {
        // `async` keyword inside loop implicitly creates a Promise
        // (though for heavier tasks, worker_threads would be preferred)
        promises.push((async () => {
            try {
                const res = await fibonacci(i);
                // console.log(`JS: Fib(${i}) = ${res}`); // uncomment for detailed output
                return res;
            } catch (e) {
                console.error(`JavaScript: Fib(${i}) generated an exception: ${e.message}`);
                return -1; // Indicate error
            }
        })());
    }

    // Wait for all promises to resolve
    await Promise.all(promises);

    const endTime = performance.now();
    console.log(`JavaScript: Total time taken for Fibonacci up to ${maxN}: ${(endTime - startTime) / 1000.0:.4f} seconds`);
    console.log("JavaScript calculation complete.\n");
}

// Example call: calculateConcurrentFibonacci(15, 4);
