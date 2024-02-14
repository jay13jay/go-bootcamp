package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	firstFour := make([]int, 1, 4)
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	_ = copy(firstFour, nums)

	// s.Show("firstFour", firstFour)
	rainChance()
}

func rainChance() {
	data := []float64{10, 25, 30, 50}
	s.Show("Probabilities", data)

	newData := []float64{99, 100}

	copy(data, newData)
	s.Show("New data", data)

	var probTot float64
	for _, p := range data {
		probTot += p
	}
	avg := probTot / float64(len(data))

	fmt.Printf("Average chance of rain: %.f%%\n", avg)
}
