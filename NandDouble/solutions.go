package main

import "fmt"

func checkIfExist(arr []int) bool {
	seen := make(map[int]int, len(arr))

	for _, v := range arr {
		_, double := seen[2*v]

		var half bool

		if v%2 == 0 {
			_, half = seen[v/2]
		}

		if double || half {
			return true
		} else {
			seen[v] = v
		}
	}

	return false
}
func main() {
	arr := []int{-10, 12, -20, -8, 15}

	fmt.Println(checkIfExist(arr))
}
