package main

// simpleArraySum returns the sum of an array of integers
func simpleArraySum(ar []int32) int32 {
	var total int32 = 0
	for _, val := range ar {
		total += val
	}

	return total
}
