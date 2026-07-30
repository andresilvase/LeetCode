package main

func StrictlyIncreasingArrayLength(nums []int) int {

	longest := 1
	current := 1

	if len(nums) == 0 {
		return 0
	}

	for i := 1; i < len(nums); i++ {

		if nums[i] > nums[i-1] {
			current++
		} else {
			current = 1
		}

		longest = max(longest, current)
	}

	return longest
}
