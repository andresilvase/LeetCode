package main

type TestCase struct {
	name     string
	input    []int
	expected int
}

var increasingTests = []TestCase{
	{
		name:     "empty array",
		input:    []int{},
		expected: 0,
	},
	{
		name:     "single element",
		input:    []int{5},
		expected: 1,
	},
	{
		name:     "two increasing elements",
		input:    []int{1, 2},
		expected: 2,
	},
	{
		name:     "two decreasing elements",
		input:    []int{2, 1},
		expected: 1,
	},
	{
		name:     "two equal elements",
		input:    []int{5, 5},
		expected: 1,
	},
	{
		name:     "fully increasing",
		input:    []int{1, 2, 3, 4, 5},
		expected: 5,
	},
	{
		name:     "fully decreasing",
		input:    []int{5, 4, 3, 2, 1},
		expected: 1,
	},
	{
		name:     "all elements equal",
		input:    []int{7, 7, 7, 7, 7},
		expected: 1,
	},
	{
		name:     "longest sequence at beginning",
		input:    []int{1, 2, 3, 4, 0, -1},
		expected: 4,
	},
	{
		name:     "longest sequence in middle",
		input:    []int{9, 1, 2, 3, 4, 0},
		expected: 4,
	},
	{
		name:     "longest sequence at end",
		input:    []int{9, 8, 1, 2, 3, 4},
		expected: 4,
	},
	{
		name:     "duplicate breaks sequence",
		input:    []int{1, 2, 3, 3, 4, 5},
		expected: 3,
	},
	{
		name:     "multiple duplicates",
		input:    []int{1, 1, 2, 2, 3, 3},
		expected: 2,
	},
	{
		name:     "multiple increasing sequences",
		input:    []int{1, 2, 0, 3, 4, 5, 1, 2},
		expected: 4,
	},
	{
		name:     "multiple sequences with same length",
		input:    []int{1, 2, 0, 3, 4, 0, 5, 6},
		expected: 3,
	},
	{
		name:     "alternating values",
		input:    []int{1, 3, 2, 4, 3, 5, 4, 6},
		expected: 2,
	},
	{
		name:     "negative increasing values",
		input:    []int{-5, -4, -3, -2, -1},
		expected: 5,
	},
	{
		name:     "negative decreasing values",
		input:    []int{-1, -2, -3, -4},
		expected: 1,
	},
	{
		name:     "mixed negative and positive",
		input:    []int{-3, -2, -1, 0, 1},
		expected: 5,
	},
	{
		name:     "sequence crosses zero",
		input:    []int{-2, -1, 0, 1, -5},
		expected: 4,
	},
	{
		name:     "large jumps still count",
		input:    []int{1, 100, 1000, 10000},
		expected: 4,
	},
	{
		name:     "small decreases break sequence",
		input:    []int{1, 2, 3, 2, 3, 4, 5},
		expected: 4,
	},
	{
		name:     "repeated local peaks",
		input:    []int{1, 2, 1, 2, 1, 2},
		expected: 2,
	},
	{
		name:     "long sequence after equal values",
		input:    []int{3, 3, 3, 1, 2, 3, 4},
		expected: 4,
	},
	{
		name:     "long sequence before equal values",
		input:    []int{1, 2, 3, 4, 4, 4},
		expected: 4,
	},
	{
		name:     "minimum int-like values",
		input:    []int{-1_000_000_000, -999_999_999, 0},
		expected: 3,
	},
	{
		name:     "maximum int-like values",
		input:    []int{999_999_998, 999_999_999, 1_000_000_000},
		expected: 3,
	},
	{
		name:     "large fully increasing array",
		input:    increasingArray(10_000),
		expected: 10_000,
	},
	{
		name:     "large fully decreasing array",
		input:    decreasingArray(10_000),
		expected: 1,
	},
	{
		name:     "large constant array",
		input:    constantArray(10_000, 42),
		expected: 1,
	},
	{
		name:     "large alternating array",
		input:    alternatingArray(10_000),
		expected: 2,
	},
}

var decreasingTests = []TestCase{
	{
		name:     "empty array",
		input:    []int{},
		expected: 0,
	},
	{
		name:     "single element",
		input:    []int{5},
		expected: 1,
	},
	{
		name:     "two increasing elements",
		input:    []int{1, 2},
		expected: 1,
	},
	{
		name:     "two decreasing elements",
		input:    []int{2, 1},
		expected: 2,
	},
	{
		name:     "two equal elements",
		input:    []int{5, 5},
		expected: 1,
	},
	{
		name:     "fully increasing",
		input:    []int{1, 2, 3, 4, 5},
		expected: 1,
	},
	{
		name:     "fully decreasing",
		input:    []int{5, 4, 3, 2, 1},
		expected: 5,
	},
	{
		name:     "all elements equal",
		input:    []int{7, 7, 7, 7, 7},
		expected: 1,
	},
	{
		name:     "longest sequence at beginning",
		input:    []int{1, 2, 3, 4, 0, -1},
		expected: 3,
	},
	{
		name:     "longest sequence in middle",
		input:    []int{9, 1, 2, 3, 4, 0},
		expected: 2,
	},
	{
		name:     "longest sequence at end",
		input:    []int{9, 8, 1, 2, 3, 4},
		expected: 3,
	},
	{
		name:     "duplicate breaks sequence",
		input:    []int{1, 2, 3, 3, 4, 5},
		expected: 1,
	},
	{
		name:     "multiple duplicates",
		input:    []int{1, 1, 2, 2, 3, 3},
		expected: 1,
	},
	{
		name:     "multiple monotonic sequences",
		input:    []int{1, 2, 0, 3, 4, 5, 1, 2},
		expected: 2,
	},
	{
		name:     "multiple sequences with same length",
		input:    []int{1, 2, 0, 3, 4, 0, 5, 6},
		expected: 2,
	},
	{
		name:     "alternating values",
		input:    []int{1, 3, 2, 4, 3, 5, 4, 6},
		expected: 2,
	},
	{
		name:     "negative increasing values",
		input:    []int{-5, -4, -3, -2, -1},
		expected: 1,
	},
	{
		name:     "negative decreasing values",
		input:    []int{-1, -2, -3, -4},
		expected: 4,
	},
	{
		name:     "mixed negative and positive",
		input:    []int{-3, -2, -1, 0, 1},
		expected: 1,
	},
	{
		name:     "sequence crosses zero",
		input:    []int{-2, -1, 0, 1, -5},
		expected: 2,
	},
	{
		name:     "large jumps still count",
		input:    []int{1, 100, 1000, 10000},
		expected: 1,
	},
	{
		name:     "small decrease breaks sequence",
		input:    []int{1, 2, 3, 2, 3, 4, 5},
		expected: 2,
	},
	{
		name:     "repeated local peaks",
		input:    []int{1, 2, 1, 2, 1, 2},
		expected: 2,
	},
	{
		name:     "long sequence after equal values",
		input:    []int{3, 3, 3, 1, 2, 3, 4},
		expected: 2,
	},
	{
		name:     "long sequence before equal values",
		input:    []int{1, 2, 3, 4, 4, 4},
		expected: 1,
	},
	{
		name:     "minimum int-like values",
		input:    []int{-1_000_000_000, -999_999_999, 0},
		expected: 1,
	},
	{
		name:     "maximum int-like values",
		input:    []int{999_999_998, 999_999_999, 1_000_000_000},
		expected: 1,
	},
	{
		name:     "increasing then decreasing",
		input:    []int{1, 2, 3, 4, 3, 2, 1},
		expected: 4,
	},
	{
		name:     "decreasing then increasing",
		input:    []int{4, 3, 2, 1, 2, 3, 4, 5},
		expected: 4,
	},
	{
		name:     "longer decreasing sequence",
		input:    []int{1, 2, 8, 7, 6, 5, 4},
		expected: 5,
	},
	{
		name:     "longer increasing sequence",
		input:    []int{9, 8, 1, 2, 3, 4, 5},
		expected: 3,
	},
	{
		name:     "equal increasing and decreasing lengths",
		input:    []int{1, 2, 3, 2, 1},
		expected: 3,
	},
	{
		name:     "mountain shape",
		input:    []int{1, 3, 5, 7, 6, 4, 2},
		expected: 4,
	},
	{
		name:     "valley shape",
		input:    []int{7, 5, 3, 1, 2, 4, 6},
		expected: 4,
	},
	{
		name:     "large fully increasing array",
		input:    increasingArray(10_000),
		expected: 1,
	},
	{
		name:     "large fully decreasing array",
		input:    decreasingArray(10_000),
		expected: 10_000,
	},
	{
		name:     "large constant array",
		input:    constantArray(10_000, 42),
		expected: 1,
	},
	{
		name:     "large alternating array",
		input:    alternatingArray(10_000),
		expected: 2,
	},
}

var incOrDecTests = []TestCase{
	{
		name:     "empty array",
		input:    []int{},
		expected: 0,
	},
	{
		name:     "single element",
		input:    []int{5},
		expected: 1,
	},
	{
		name:     "two increasing elements",
		input:    []int{1, 2},
		expected: 2,
	},
	{
		name:     "two decreasing elements",
		input:    []int{2, 1},
		expected: 2,
	},
	{
		name:     "two equal elements",
		input:    []int{5, 5},
		expected: 1,
	},
	{
		name:     "fully increasing",
		input:    []int{1, 2, 3, 4, 5},
		expected: 5,
	},
	{
		name:     "fully decreasing",
		input:    []int{5, 4, 3, 2, 1},
		expected: 5,
	},
	{
		name:     "all elements equal",
		input:    []int{7, 7, 7, 7, 7},
		expected: 1,
	},
	{
		name:     "longest sequence at beginning",
		input:    []int{1, 2, 3, 4, 0, -1},
		expected: 4,
	},
	{
		name:     "longest sequence in middle",
		input:    []int{9, 1, 2, 3, 4, 0},
		expected: 4,
	},
	{
		name:     "longest sequence at end",
		input:    []int{9, 8, 1, 2, 3, 4},
		expected: 4,
	},
	{
		name:     "duplicate breaks sequence",
		input:    []int{1, 2, 3, 3, 4, 5},
		expected: 3,
	},
	{
		name:     "multiple duplicates",
		input:    []int{1, 1, 2, 2, 3, 3},
		expected: 2,
	},
	{
		name:     "multiple monotonic sequences",
		input:    []int{1, 2, 0, 3, 4, 5, 1, 2},
		expected: 4,
	},
	{
		name:     "multiple sequences with same length",
		input:    []int{1, 2, 0, 3, 4, 0, 5, 6},
		expected: 3,
	},
	{
		name:     "alternating values",
		input:    []int{1, 3, 2, 4, 3, 5, 4, 6},
		expected: 2,
	},
	{
		name:     "negative increasing values",
		input:    []int{-5, -4, -3, -2, -1},
		expected: 5,
	},
	{
		name:     "negative decreasing values",
		input:    []int{-1, -2, -3, -4},
		expected: 4,
	},
	{
		name:     "mixed negative and positive",
		input:    []int{-3, -2, -1, 0, 1},
		expected: 5,
	},
	{
		name:     "sequence crosses zero",
		input:    []int{-2, -1, 0, 1, -5},
		expected: 4,
	},
	{
		name:     "large jumps still count",
		input:    []int{1, 100, 1000, 10000},
		expected: 4,
	},
	{
		name:     "small decrease breaks sequence",
		input:    []int{1, 2, 3, 2, 3, 4, 5},
		expected: 4,
	},
	{
		name:     "repeated local peaks",
		input:    []int{1, 2, 1, 2, 1, 2},
		expected: 2,
	},
	{
		name:     "long sequence after equal values",
		input:    []int{3, 3, 3, 1, 2, 3, 4},
		expected: 4,
	},
	{
		name:     "long sequence before equal values",
		input:    []int{1, 2, 3, 4, 4, 4},
		expected: 4,
	},
	{
		name:     "minimum int-like values",
		input:    []int{-1_000_000_000, -999_999_999, 0},
		expected: 3,
	},
	{
		name:     "maximum int-like values",
		input:    []int{999_999_998, 999_999_999, 1_000_000_000},
		expected: 3,
	},
	{
		name:     "increasing then decreasing",
		input:    []int{1, 2, 3, 4, 3, 2, 1},
		expected: 4,
	},
	{
		name:     "decreasing then increasing",
		input:    []int{4, 3, 2, 1, 2, 3, 4, 5},
		expected: 5,
	},
	{
		name:     "longer decreasing sequence",
		input:    []int{1, 2, 8, 7, 6, 5, 4},
		expected: 5,
	},
	{
		name:     "longer increasing sequence",
		input:    []int{9, 8, 1, 2, 3, 4, 5},
		expected: 5,
	},
	{
		name:     "equal increasing and decreasing lengths",
		input:    []int{1, 2, 3, 2, 1},
		expected: 3,
	},
	{
		name:     "mountain shape",
		input:    []int{1, 3, 5, 7, 6, 4, 2},
		expected: 4,
	},
	{
		name:     "valley shape",
		input:    []int{7, 5, 3, 1, 2, 4, 6},
		expected: 4,
	},
	{
		name:     "large fully increasing array",
		input:    increasingArray(10_000),
		expected: 10_000,
	},
	{
		name:     "large fully decreasing array",
		input:    decreasingArray(10_000),
		expected: 10_000,
	},
	{
		name:     "large constant array",
		input:    constantArray(10_000, 42),
		expected: 1,
	},
	{
		name:     "large alternating array",
		input:    alternatingArray(10_000),
		expected: 2,
	},
}

func increasingArray(size int) []int {
	array := make([]int, size)

	for index := range array {
		array[index] = index
	}

	return array
}

func decreasingArray(size int) []int {
	array := make([]int, size)

	for index := range array {
		array[index] = size - index
	}

	return array
}

func preview(array []int) []int {
	const limit = 20

	if len(array) <= limit {
		return array
	}

	return array[:limit]
}

func constantArray(size int, value int) []int {
	result := make([]int, size)

	for i := range result {
		result[i] = value
	}

	return result
}

func alternatingArray(size int) []int {
	result := make([]int, size)

	for i := range result {
		if i%2 == 0 {
			result[i] = 0
		} else {
			result[i] = 1
		}
	}

	return result
}
