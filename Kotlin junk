import kotlinx.coroutines.*
import java.util.concurrent.ConcurrentHashMap
import kotlin.system.measureTimeMillis

// Thread-safe memoization cache
val fibCacheKotlin = ConcurrentHashMap<Int, Long>()

// Simulate asynchronous work using delay
suspend fun simulateWorkKotlin(ms: Int) {
    delay(ms.toLong())
}

// Suspend function for Fibonacci calculation with memoization
suspend fun fibonacciKotlin(n: Int): Long =
    coroutineScope { // Provides structured concurrency context
        if (n < 0) throw IllegalArgumentException("Input must be a non-negative integer.")
        if (n == 0) return@coroutineScope 0
        if (n == 1) return@coroutineScope 1

        // Check cache (needs synchronized access or ConcurrentHashMap)
        // Note: ConcurrentHashMap is thread-safe.
        fibCacheKotlin[n]?.let {
            return@coroutineScope it // Return cached value
        }

        simulateWorkKotlin(1) // Simulate work

        // Launch concurrent tasks for sub-problems
        val deferred1 = async(Dispatchers.Default) { fibonacciKotlin(n - 1) }
        val deferred2 = async(Dispatchers.Default) { fibonacciKotlin(n - 2) }

        // Await results
        val result1 = deferred1.await()
        val deferred2.await()

        val result = result1 + result2

        fibCacheKotlin[n] = result // Store in cache
        result
    }

// Function to calculate Fibonacci numbers concurrently
suspend fun calculateConcurrentFibonacciKotlin(maxN: Int, numThreads: Int) {
    println("\n--- Kotlin Example ---")
    println("Kotlin: Calculating Fibonacci numbers up to $maxN concurrently using Coroutines...")

    val totalTime = measureTimeMillis {
        // Clear cache for a fresh run
        fibCacheKotlin.clear()
        
        // Use coroutineScope to ensure all launched coroutines complete
        coroutineScope {
            val jobs = List(maxN + 1) { i ->
                launch(Dispatchers.Default) { // Use Dispatchers.Default for CPU-bound work
                    try {
                        val res = fibonacciKotlin(i)
                        // println("Kotlin: Fib($i) = $res") // uncomment for verbose output
                    } catch (e: Exception) {
                        System.err.println("Kotlin: Error calculating Fib($i): ${e.message}")
                    }
                }
            }
            // jobs.joinAll() // This is implicitly handled by the outer coroutineScope if launched within it
        }
    }

    println("Kotlin: Total time taken for Fibonacci up to $maxN: ${totalTime / 1000.0} seconds")
    println("Kotlin calculation complete.\n")
}

// Example usage (requires kotlinx-coroutines-core library):
// To run:
// fun main() = runBlocking {
//     calculateConcurrentFibonacciKotlin(15, 4)
// }
