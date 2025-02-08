package main

import "fmt"

func main() {
	countApplesAndOranges(7, 10, 4, 12, []int32{2, 3, -4}, []int32{3, -2, -4})
}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var (
		totalApples  int32
		totalOranges int32
	)

	for _, apple := range apples {
		totalDistance := apple + a
		if totalDistance >= s && totalDistance <= t {
			totalApples++
		}
	}

	for _, orange := range oranges {
		totalDistance := orange + b
		if totalDistance >= s && totalDistance >= t {
			totalOranges++
		}
	}

	fmt.Println(totalApples)
	fmt.Println(totalOranges)
}
