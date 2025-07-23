defmodule FibCalculator do
  # For simplicity, we'll use a GenServer to manage the cache.
  # A simple Map could be used, but a GenServer explicitly manages state concurrently.
  use GenServer

  # Client API
  def start_link(_opts) do
    GenServer.start_link(__MODULE__, %{})
  end

  def calculate(pid, n) do
    GenServer.call(pid, {:calculate, n})
  end

  # GenServer Callbacks
  @impl true
  def init(initial_state) do
    {:ok, initial_state}
  end

  # Handler for :calculate requests
  @impl true
  def handle_call({:calculate, n}, _from) do
    # Use Task.async for lightweight concurrency. The result is managed by the calling process (this GenServer).
    # We simulate work using a sleep equivalent in Elixir.
    task = Task.async(fn -> fibonacci_elixir_recursive(n) end)
    result = Task.await(task) # Blocks the GenServer *call* temporarily for this calculation
    
    # Storing result would typically happen here if cache was managed by GenServer.
    # For demonstration, we directly return the computed value from Task.async.
    # A more typical Cache GenServer would do `{:reply, result, %{n => result} = state}`.
    # Since this isn't a persistent cache within the GenServer here, we just return.

    # If using a shared cache managed externally, Task.async allows returning result to be picked up.
    {:reply, result, %{}} # Returning empty map state as cache is not actively managed here directly by GenServer
  end

  # Simulate work (simple sleep, blocking for simplicity in example)
  defp simulate_work_elixir(ms) do
    Process.sleep(ms)
  end
  
  # Recursive Fibonacci, with basic memoization if adapted to a stateful process
  defp fibonacci_elixir_recursive(n) when is_integer(n) and n >= 0 do
    simulate_work_elixir(1) # Simulate work
    cond do
      n == 0 -> 0
      n == 1 -> 1
      true ->
        # Here, to make memoization work, we'd need a shared ETS table or a stateful GenServer
        # that we query recursively. For Task.async demonstration, we compute without it here.
        # The Task.async/await manages the computation, not shared cache coordination itself.
        
        # A proper cached version:
        # Cache = :ets.new(:fib_cache_ets, [:named_table, read_concurrency: true])
        # case :ets.lookup_element(Cache, n, 1) do
        #   nil ->
        #     result = fibonacci_elixir_recursive(n-1) + fibonacci_elixir_recursive(n-2)
        #     :ets.insert(Cache, {n, result})
        #     result
        #   val -> val
        # end
        
        # Without ETS cache lookup within recursive step (for simple Task demo):
        left = spawn(fn -> fibonacci_elixir_elixir_recursive(n - 1) end) # Using spawn for illustration, not ideal. Task is better.
        right = spawn(fn -> fibonacci_elixir_elixir_recursive(n - 2) end)
        
        # Proper Elixir would use Task.async within this recursion as well or pass message passing.
        # Let's simplify this to demonstrate Task.async more directly on top-level calls.
        
        # Correct use of Task for concurrent sub-problems if not using process state:
        task_left = Task.async(fn -> fibonacci_elixir_recursive(n - 1) end)
        task_right = Task.async(fn -> fibonacci_elixir_recursive(n - 2) end)

        Task.await(task_left) + Task.await(task_right)
    end
  end
end

defmodule Runner do
  # This runner will orchestrate multiple concurrent Fibonacci calculations.
  # For proper memoization across concurrent calls, ETS or a dedicated cache process is ideal.
  # This example uses Task.async to highlight concurrent execution of independent calculations.
  
  def run_fib_concurrent(max_n, num_workers) do
    # For demonstration, we won't create a persistent cache managed by Elixir processes here
    # but focus on demonstrating concurrency via Tasks for calculating Fib values.
    
    # Starting the calculator process (GenServer) if cache was managed there
    # {:ok, calculator_pid} = FibCalculator.start_link([])
    
    start_time = :os.system_time(:milli_seconds)

    # Use Task.Supervisor to manage concurrent tasks robustly
    tasks = for i <- 0..max_n do
      Task.async(fn -> 
        # The FibCalculator.calculate would call the recursive fibonacci_elixir_recursive.
        # We are bypassing direct GenServer call for simplicity of fibonacci computation demo
        # and just call the simulated fib_elixir function.
        # Task.async(fn -> FibCalculator.calculate(calculator_pid, i) end) # If using GenServer
        Task.async(fn -> FibonacciCalculator.fibonacci_elixir_recursive(i) end)
      end)
    end

    results = Enum.map(tasks, fn task -> Task.await(task) end)
    
    end_time = :os.system_time(:milli_seconds)
    elapsed = (end_time - start_time) / 1000.0

    IO.puts("\n--- Elixir Example ---")
    IO.puts("Elixir: Calculating Fibonacci numbers up to #{max_n} concurrently using Tasks...")
    # IO.inspect(results, label: "Elixir Results") # Uncomment for verbose results
    IO.puts("Elixir: Total time taken for Fibonacci up to #{max_n}: #{elapsed} seconds")
    IO.puts("Elixir calculation complete.\n")
  end
end

# To run:
# ExSg = IO.read(:stdio, :line) |> String.trim() # for prompt or set directly
# maxN = ExSg |> Integer.parse() |> elem(0) 
# Runner.run_fib_concurrent(max_n, 4) # Example: max_n = 15
