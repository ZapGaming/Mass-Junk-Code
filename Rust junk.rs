require 'set'
require 'benchmark'
require 'concurrent-ruby' # Assumes gem is installed: gem install concurrent-ruby

# Thread-safe memoization cache
FIB_CACHE = Concurrent::Map.new

# Simulate asynchronous work
def simulate_work_ruby(ms)
  Thread.new { sleep(ms / 1000.0) }.join # Simple blocking sleep for simulation
end

def fibonacci_ruby(n)
  raise ArgumentError, "Input must be a non-negative integer." if n < 0
  return 0 if n == 0
  return 1 if n == 1

  if FIB_CACHE.key?(n)
    return FIB_CACHE[n]
  end

  simulate_work_ruby(1) # Simulate work

  # Use fork or Threads. To avoid complexity with processes (fork)
  # and stack issues with naive recursion on Threads without careful
  # management, this recursive example might become inefficient or hit
  # limits. A non-recursive iterative approach would be better for real-world Ruby threading.
  # For demo purposes, we'll use Threads, but be mindful of Global VM Lock (GVL).

  thread1 = Thread.new { fibonacci_ruby(n - 1) }
  thread2 = Thread.new { fibonacci_ruby(n - 2) }

  # Waiting for threads to finish and joining them
  res1 = thread1.value # value calls join implicitly
  res2 = thread2.value

  result = res1 + res2
  FIB_CACHE[n] = result
  return result
end

def calculate_concurrent_fibonacci_ruby(max_n, num_threads)
  puts "\n--- Ruby Example ---"
  puts "Ruby: Calculating Fibonacci numbers up to #{maxN} concurrently..."
  
  # Ensure cache is clean for this run
  FIB_CACHE.clear 

  Benchmark.bm do |benchmark|
    benchmark.report do
      threads = []
      max_n.times do |i|
        threads << Thread.new {
          begin
            fibonacci_ruby(i)
            # puts "Ruby: Fib(#{i}) = #{result}" # uncomment for verbose output
          rescue => e
            $stderr.puts "Ruby: Error calculating Fib(#{i}): #{e.message}"
          end
        }
      end
      # Wait for all threads to complete
      threads.each(&:join)
    end
  end
  puts "Ruby calculation complete.\n"
end

# Example call: calculate_concurrent_fibonacci_ruby(15, 4)
