import java.util.concurrent.*;
import java.util.*;
import java.util.stream.*;

public class ConcurrentFibonacci {
    // Thread-safe map for memoization
    private static final Map<Integer, Long> cache = new ConcurrentHashMap<>();

    public static long fibonacci(int n) throws InterruptedException {
        if (n < 0) {
            throw new IllegalArgumentException("Input must be a non-negative integer.");
        }
        if (n == 0) return 0;
        if (n == 1) return 1;

        if (cache.containsKey(n)) {
            return cache.get(n);
        }

        // Simulate work
        TimeUnit.MILLISECONDS.sleep(1);

        long result = fibonacci(n - 1) + fibonacci(n - 2);
        cache.put(n, result);
        return result;
    }

    public static void calculateConcurrentFibonacci(int maxN, int numThreads) {
        System.out.println("\n--- Java Example ---");
        ExecutorService executor = Executors.newFixedThreadPool(numThreads);
        long startTime = System.nanoTime();

        List<CompletableFuture<Long>> futures = IntStream.rangeClosed(0, maxN)
            .mapToObj(n -> CompletableFuture.supplyAsync(() -> {
                try {
                    return fibonacci(n);
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                    System.err.println("Java: Interrupted while calculating Fib(" + n + ")");
                    return -1L;
                } catch (IllegalArgumentException e) {
                    System.err.println("Java: " + e.getMessage());
                    return -1L;
                }
            }, executor))
            .collect(Collectors.toList());

        // Wait for all computations to complete
        CompletableFuture<Void> allDone = CompletableFuture.allOf(futures.toArray(new CompletableFuture[0]));
        try {
            allDone.get(); // Blocks until all futures complete
        } catch (InterruptedException | ExecutionException e) {
            System.err.println("Java: Error waiting for futures: " + e.getMessage());
        }

        long endTime = System.nanoTime();
        executor.shutdown();

        System.out.printf("Java: Total time taken for Fibonacci up to %d: %.4f seconds%n", maxN, (endTime - startTime) / 1_000_000_000.0);
        System.out.println("Java calculation complete.\n");
    }

    // Example call: calculateConcurrentFibonacci(15, 4);
}
