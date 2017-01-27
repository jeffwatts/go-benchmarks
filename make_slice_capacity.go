package go_benchmarks

type TextRange struct {
	matchingRange []int
	text          string
}

func MakeWithCapacity(numToFill int, text string, capacity int) []TextRange {
	result := make([]TextRange, 0, capacity)

	for numFilled := 0; numFilled < numToFill; numFilled++ {
		result = append(result, TextRange{[]int{1,2}, text})
	}

	return result
}

func MakeWithoutCapacity(numToFill int, text string) []TextRange {
	result := make([]TextRange, 0)

	for numFilled := 0; numFilled < numToFill; numFilled++ {
		result = append(result, TextRange{[]int{1,2}, text})
	}

	return result
}
