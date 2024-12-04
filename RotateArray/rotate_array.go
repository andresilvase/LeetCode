package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)

	if n <= 1 || k == 0 {
		return
	}

	k %= n

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 2

	fmt.Printf("Original Array %v\n", nums)
	rotate(nums, k)
	fmt.Printf("Array rotated by %d positions %v\n", k, nums)
}
