package main

func StrictlyIncreasingArrayLength(nums []int) int {

	var globalCheck = 1
	var localCheck = 1

	for i := range nums {
		j := i + 1

		if j == len(nums) {
			break
		}

		// fmt.Printf("nums[%d]: %d - nums[%d]: %d | localCheck: %d\n", i, nums[i], j, nums[j], localCheck)

		if nums[j] > nums[i] {
			localCheck++
		} else {
			globalCheck = max(globalCheck, localCheck)
			localCheck = 0
		}
	}

	return max(globalCheck, localCheck)
}
