package main

import "fmt"

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

func StrictlyDecreasingArrayLength(nums []int) int {

	longest := 1
	current := 1

	if len(nums) == 0 {
		return 0
	}

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
	longest := 1
	shortest := 1
	currentLongest := 1
	currentShortest := 1

	if len(nums) == 0 {
		return 0
	}

	for i := 1; i < len(nums); i++ {

		if nums[i] > nums[i-1] {
			currentLongest++
			currentShortest = 1
		} else if nums[i] < nums[i-1] {
			currentShortest++
			currentLongest = 1
		} else {
			currentLongest = 1
			currentShortest = 1
		}

		shortest = max(shortest, currentShortest)
		longest = max(longest, currentLongest)
	}

	return max(longest, shortest)
}

func submitCodeFor(challengeName string, testSet []TestCase, solution func([]int) int) {
	fmt.Printf("\nRunning tests for %s...\n", challengeName)

	for i, test := range testSet {
		result := solution(test.input)

		if result == test.expected {
			fmt.Printf("Test %d:- PASS\n", i+1)
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
