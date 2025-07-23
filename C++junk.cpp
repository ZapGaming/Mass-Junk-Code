#include <iostream>
#include <vector>
#include <thread>
#include <future>
#include <unordered_map>
#include <mutex>
#include <chrono>
#include <functional>

// Thread-safe memoization cache
std::unordered_map<int, long long> fib_cache;
std::mutex cache_mutex;

long long fibonacci_cpp(int n) {
    if (n < 0) {
        throw std::invalid_argument("Input must be a non-negative integer.");
    }
    if (n == 0) return 0;
    if (n == 1) return 1;

    // Lock the mutex to access the cache
    std::lock_guard<std::mutex> lock(cache_mutex);
    if (fib_cache.count(n)) {
        return fib_cache[n];
    }
    
    // Unlock mutex before recursive calls to avoid deadlock if recursive calls lock again.
    // This means cache isn't protected during the "simulate work", but subsequent lookups will be safe.
    // A more robust approach might involve atomic checks or separate read/write locks if performance critical.
    lock.unlock(); 

    // Simulate work (this happens outside the mutex guard to allow concurrent computation)
    std::this_thread::sleep_for(std::chrono::milliseconds(1)); 

    // Compute recursively, deferring to std::async can help manage threads
    // Note: Launching async here for *each* recursive call adds overhead. 
    // For true complex async Fibonacci, you'd manage a thread pool.
    // Here, we use std::async mainly to showcase potential parallelism and future management.
    auto fut1 = std::async(std::launch::async, fibonacci_cpp, n - 1);
    auto fut2 = std::async(std::launch::async, fibonacci_cpp, n - 2);
    
    long long result = fut1.get() + fut2.get();

    // Re-acquire lock to store the result
    lock_guard<std::mutex> store_lock(cache_mutex);
    fib_cache[n] = result;
    return result;
}

void calculateConcurrentFibonacci_cpp(int maxN, int numThreads) {
    std::cout << "\n--- C++ Example ---" << std::endl;
    std::cout << "C++: Calculating Fibonacci numbers up to " << maxN << " concurrently using " << numThreads << " threads (via std::async)..." << std::endl;
    auto start_time = std::chrono::high_resolution_clock::now();

    std::vector<std::future<long long>> futures;
    for (int i = 0; i <= maxN; ++i) {
        // std::launch::async ensures a new thread or task from a pool is used.
        // std::ref() is used because async copies arguments by default.
        futures.push_back(std::async(std::launch::async, fibonacci_cpp, i));
    }

    std::vector<long long> results;
    results.reserve(maxN + 1);
    for (auto& fut : futures) {
        try {
            results.push_back(fut.get());
            // std::cout << "C++: Fib(" << results.size() -1 << ") = " << results.back() << std::endl; // optional verbose output
        } catch (const std::exception& e) {
            std::cerr << "C++: Exception caught: " << e.what() << std::endl;
            results.push_back(-1); // Indicate error
        }
    }

    auto end_time = std::chrono::high_resolution_clock::now();
    std::chrono::duration<double> elapsed = end_time - start_time;
    std::cout << "C++: Total time taken for Fibonacci up to " << maxN << ": " << elapsed.count() << " seconds" << std::endl;
    std::cout << "C++ calculation complete.\n" << std::endl;
}

// Example call: calculateConcurrentFibonacci_cpp(15, 4);
