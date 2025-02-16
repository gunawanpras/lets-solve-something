package main

import (
	"sort"
)

func main() {
	println(migratoryBirds([]int32{1, 4, 4, 4, 5, 3}))
	println(migratoryBirds([]int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}))
}

type indexedValue struct {
	index int
	value int
}

func migratoryBirds(arr []int32) int32 {
	birdTypes := make([]indexedValue, 5)
	for _, v := range arr {
		birdTypes[int(v-1)].index = int(v)
		birdTypes[int(v-1)].value++
	}

	sort.Slice(birdTypes, func(i, j int) bool {
		return birdTypes[i].value > birdTypes[j].value
	})

	var (
		typeId    int
		typeValue int
	)

	for _, v := range birdTypes {
		if typeId <= 0 && typeValue <= 0 {
			typeId = v.index
			typeValue = v.value
			continue
		}

		if v.index < typeId && v.value == typeValue {
			typeId = v.index
			typeValue = v.value
		}
	}

	return int32(typeId)
}
