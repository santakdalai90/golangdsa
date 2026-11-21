- # Dynamic Programming
	- Dynamic programming is a class of problems which has these properties
		- 1. **Optimal sub-structure:** Problem can be solved using recursion where a solution at a stage can be aggregated from solutions of its substages.
		- 2. **Overlapping sub-problems:** Number of unique nodes in the recursion tree is lesser than the total number of nodes in it.
	- ## Solving a dynamic programming question
		- A dynamic programming question can be solved if you can figure out these for the given question
			1. Recurrence Relationship
			2. Number of distinct states

	- ## Approaches for solving dynamic problem
		- ### Top Down approch
			- Recursion and memoization
		- ### Bottom Up approach
			- Iterative approach
			- Start building your answer from the base case
	- ## Problems
		- ### Problem 1: Nth Fibonacci number
		    
		  Recurrence Relationship: fib(n) = fib(n-1) + fib(n-2)  
		    
		  Number of distinct states: We can see clearly for calculating fib(n) we'd need to know fib(n-1), fib(n-2), fib(n-3), ...,  
		  fib(0). So we have n distinct states. We can memoize this in an array of size n.  
		    
		  ```go
		  func fib(n int, memo []int) int {
		  	if memo[n] != -1 {
		  		return memo[n]
		  	}
		  	memo[n] = fib(n-1, memo) + fib(n-2, memo)
		  	return memo[n]
		  }
		  ```
		- ### Problem 2: Number of ways in a staircase to reach a level
			- In a given staircase you're only allowed to make moves where you can jump up 1 step or you can jump up 2 steps. So to reach a given step how many different ways exist?
			- Base Cases
				- Only 1 step. Number of ways = 1
				- 2 steps. Number of ways = 2, i.e.,
					- first, 0 -> 1 and then 1 -> 2
					- second, directly from 0 -> 2
				- 3 steps. Number of ways = 3, i.e,
					- first, 0 ->1, 1 -> 2, 2 -> 3
					- second, 0 -> 2, 2 -> 3
					- third, 0 -> 1, 1 -> 3
					- Here you may observe that Number of ways for step 3 = number of ways for step 1 + number of ways for step 2
			- So for n steps, Number of ways = number of ways for step n-1 + number of ways for step n-2
			- In other words we can say that the number of ways is equal to Fib(n)
		-
