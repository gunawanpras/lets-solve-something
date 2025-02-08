package main

func main() {
	// total := birthday([]int32{1, 2, 1, 3, 2}, 3, 2)
	// total := birthday([]int32{1, 1, 1, 1, 1, 1}, 3, 2)
	// total := birthday([]int32{4}, 4, 1)
	total := birthday([]int32{2, 5, 1, 3, 4, 4, 3, 5, 1, 1, 2, 1, 4, 1, 3, 3, 4, 2, 1}, 18, 7)

	println(total)
}

func birthday(s []int32, d int32, m int32) int32 {
	var (
		limit     int32
		totalBars int32
		sLength   int32 = int32(len(s))
	)

	if sLength <= 0 || d <= 0 || m <= 0 {
		return totalBars
	}

	if sLength == 1 {
		if d == s[0] {
			totalBars = 1
		}

		return totalBars

	} else {
		for i := int32(0); i <= sLength-m; i++ {
			var dividedBars int32

			if m == 1 {
				limit = m

			} else {
				limit = m + i
			}

			for j := i; j < limit; j++ {
				dividedBars += s[j]
			}

			if d == dividedBars {
				totalBars++
			}
		}

		return totalBars
	}
}
