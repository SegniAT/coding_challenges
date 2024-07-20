package main

func main() {
}

// NOT FULLY UNDERSTOOD
// explanation to the solution:
/*
Intuition
The key intuition to optimize this process lies in recognizing the pattern that the maximum sum of the array depends not just on the current partition but also on the optimal decisions made for previous partitions. DP approach will build up the solution incrementally, ensuring that at each step, we make the optimal choice based on previous computations.

Approach
-> Initialize a DP array: Initialize a DP array of length n+1 (n being the length of the input array) with all zeros. This DP array will store the maximum sum achievable up to each index. The reason for having n+1 elements is to conveniently handle the base case where no elements have been considered yet (i.e., dp[0] = 0).

-> Iterating Through the Array: For each element in the array (from the first to the last), determine the maximum sum achievable up to that point. This involves considering all possible partition sizes up to k that could end at the current element.

-> Considering All Partitions Up to Size k: For each element i, look back up to k elements to find the partition that maximizes the sum.
	This step involves two sub-steps:
	- Finding the maximum value in the current partition (max), which will be used to update all elements in the partition to this value.
	- Calculating the maximum sum for the current partition by adding the product of the maximum value (max) and the partition size (j) to the maximum sum achievable before this partition started (dp[i-j]).

-> Updating the DP Array: Update dp[i] with the maximum sum achievable by including the current element in a partition. This is done by comparing the sum calculated for each possible partition size and choosing the maximum.

Result: After filling the DP array, the maximum sum after partitioning the array is found at dp[n], which accounts for the entire array.

Complexity
Time complexity:
O(n∗k) -> n is the length of the array, and k is the maximum length of the partition.

Space complexity:
O(n) -> n is the length of the input array.
*/
func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		maximum := 0
		maxSum := 0

		for j := 1; j <= k && i-j >= 0; j++ {
			maximum = max(maximum, arr[i-j])
			maxSum = max(maxSum, dp[i-j]+maximum*j)
		}

		dp[i] = maxSum
	}

	return dp[n]
}
