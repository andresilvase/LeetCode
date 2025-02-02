package main

import (
	"fmt"
	"sort"
)

type ByStartInterval [][]int

func (interval ByStartInterval) Len() int { return len(interval) }

func (interval ByStartInterval) Swap(i, j int) {
	interval[i], interval[j] = interval[j], interval[i]
}

func (interval ByStartInterval) Less(i, j int) bool {
	return interval[i][0] < interval[j][0]
}

func merge(intervals [][]int) [][]int {
	output := make([][]int, 0)

	sort.Sort(ByStartInterval(intervals))

	for i, v := range intervals {
		if len(output) == 0 {
			output = append(output, intervals[i])
		} else {
			if intervalsHasConflicts(output, v) {
				mergedSlice := []int{output[len(output)-1][0], max(output[len(output)-1][1], v[1])}

				output[len(output)-1] = mergedSlice
			} else {
				output = append(output, intervals[i])
			}
		}
	}

	if len(output) == 0 {
		return intervals
	} else {
		return output
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func intervalsHasConflicts(output [][]int, currentInterval []int) bool {
	return currentInterval[0] <= output[len(output)-1][1]
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merge(intervals)

	fmt.Printf("Merged Intervals: %v\n", merge(intervals))
}
