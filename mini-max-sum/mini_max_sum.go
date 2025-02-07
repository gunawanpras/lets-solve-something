package main

import (
	"sort"
)

func miniMaxSum(arr []int32) (int64, int64) {
	var (
		sum  int64
		sums []int64
	)

	arrLen := len(arr)

	for i := 0; i < arrLen; i++ {
		sum = 0
		for j := 0; j < arrLen; j++ {
			if i == j {
				continue
			}

			sum += int64(arr[j])
		}

		sums = append(sums, sum)
	}

	sort.Slice(sums, func(i, j int) bool { return sums[i] < sums[j] })

	return sums[0], sums[arrLen-1]
}
