import concurrent.futures
import functools
import time

# Simple memoization decorator
def memoize(func):
    cache = {}
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        key = args
        if key not in cache:
            # Simulate some I/O-bound work or processing time
            # For this example, sleep simulates a short, blocking operation.
            time.sleep(0.0005) # 0.5 ms of simulated work per call
            cache[key] = func(*args, **kwargs)
        return cache[key]
    return wrapper

@memoize
def fibonacci(n):
    """Calculates the nth Fibonacci number recursively with memoization."""
    if n < 0:
        raise ValueError("Input must be a non-negative integer.")
    if n == 0:
        return 0
    if n == 1:
        return 1
    return fibonacci(n - 1) + fibonacci(n - 2)

def main_fib_concurrent(max_n, num_threads):
    """
    Calculates Fibonacci numbers up to max_n concurrently using a ThreadPoolExecutor
    and measures the total execution time.
    """
    print(f"Python: Calculating Fibonacci numbers up to {max_n} concurrently using {num_threads} threads...")
    start_time = time.time()
    
    results = {}
    with concurrent.futures.ThreadPoolExecutor(max_workers=num_threads) as executor:
        # Submit tasks for Fibonacci numbers from 0 to max_n
        future_to_n = {executor.submit(fibonacci, i): i for i in range(max_n + 1)}
        
        # Process results as they complete
        for future in concurrent.futures.as_completed(future_to_n):
            n = future_to_n[future]
            try:
                result = future.result()
                results[n] = result
            except Exception as exc:
                print(f'Python: Fib({n}) generated an exception: {exc}')

    end_time = time.time()
    print(f"\n--- Python Execution ---")
    print(f"Python: Total time taken for Fibonacci up to {max_n}: {end_time - start_time:.4f} seconds")
    print("Python calculation complete.\n")

# Example call: main_fib_concurrent(15, 4)
