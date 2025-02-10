package main

func main() {
	println(divisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2}))
}

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var pairs int32

	if n <= 0 || k <= 0 || ar == nil {
		return pairs
	}

	for i := int32(0); i < n; i++ {
		for j := i + 1; j <= n-1; j++ {
			if (ar[i]+ar[j])%k == 0 {
				pairs++
			}
		}
	}

	return pairs
}
