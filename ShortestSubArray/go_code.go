package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	minLength := n + 1

	dq := []int{}
	prefixSum := make([]int, n+1)

	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	for i := 0; i <= n; i++ {
		for len(dq) > 0 && prefixSum[i]-prefixSum[dq[0]] >= k {
			minLength = min(minLength, i-dq[0])
			dq = dq[1:]
		}

		for len(dq) > 0 && prefixSum[i] <= prefixSum[dq[len(dq)-1]] {
			dq = dq[:len(dq)-1]
		}

		dq = append(dq, i)
	}

	if minLength == n+1 {
		return -1
	}

	return min(minLength, n+1)
}

func main() {
	nums := []int{56, 21, 56, 35, 9}
	k := 61

	fmt.Printf("Output: %d\n", shortestSubarray(nums, k))
}
