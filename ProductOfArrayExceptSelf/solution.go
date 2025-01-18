package main

import "fmt"

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	var product = 1

	for i, v := range nums {
		answer[i] = product
		product *= v

		if i == 0 {
			continue
		} else {
			for j := i - 1; j >= 0; j-- {
				answer[j] = answer[j] * v
			}
		}
	}

	return answer
}

func main() {
	nums := []int{8, 4, 3, 2}
	fmt.Printf("Output: %v\n", productExceptSelf(nums))
}
