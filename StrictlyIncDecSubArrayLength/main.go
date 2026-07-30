package main

import "fmt"

func StrictlyIncreasingArrayLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	longest := 1
	current := 1

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

func StrictlyDecreasingArrayLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	longest := 1
	current := 1

	for i := 1; i < len(nums); i++ {

		if nums[i] < nums[i-1] {
			current++
		} else {
			current = 1
		}

		longest = max(longest, current)
	}

	return longest
}

func BothIncAndDec(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	longestInc := 1
	longestDec := 1
	currentInc := 1
	currentDec := 1

	for i := 1; i < len(nums); i++ {

		if nums[i] > nums[i-1] {
			currentInc++
			currentDec = 1
		} else if nums[i] < nums[i-1] {
			currentDec++
			currentInc = 1
		} else {
			currentInc = 1
			currentDec = 1
		}

		longestDec = max(longestDec, currentDec)
		longestInc = max(longestInc, currentInc)
	}

	return max(longestInc, longestDec)
}

func submitCodeFor(challengeName string, testSet []TestCase, solution func([]int) int) {
	fmt.Printf("\nRunning tests for %s...\n", challengeName)

	for i, test := range testSet {
		result := solution(test.input)

		if result == test.expected {
			fmt.Printf("Test %d: - PASS\n", i+1)
		} else {
			fmt.Printf(
				"Test %d: %s - FAIL | expected %d, got %d\n",
				i+1,
				test.name,
				test.expected,
				result,
			)
		}
	}
}

func main() {
	submitCodeFor("Longest Strictly Increasing Subarray", increasingTests, StrictlyIncreasingArrayLength)
	submitCodeFor("Longest Strictly Decreasing Subarray", decreasingTests, StrictlyDecreasingArrayLength)
	submitCodeFor("Longest Strictly Increasing or Decreasing Subarray", incOrDecTests, BothIncAndDec)

}
