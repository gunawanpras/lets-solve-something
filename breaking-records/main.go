package main

import "fmt"

func main() {
	result := breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1})
	fmt.Printf("RESULT: %+v\n", result)
	result = breakingRecords([]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42})
	fmt.Printf("RESULT: %+v\n", result)
}

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	var (
		min        int32
		max        int32
		mostPoint  int32
		leastPoint int32
	)

	for i, score := range scores {
		if i == 0 {
			min = score
			max = score
		}

		if score < min {
			min = score
			leastPoint++
		} else if score > max {
			max = score
			mostPoint++
		}
	}

	return []int32{mostPoint, leastPoint}
}
