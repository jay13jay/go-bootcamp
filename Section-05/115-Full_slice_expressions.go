package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true

	nums := []int{1, 3, 2, 4}

	evens := nums[2:]
	odds := append(nums[:2], 5, 7)
	// evens := append(nums[2:], 6, 8)

	s.Show("nums: ", nums)
	s.Show("odds: ", odds)
	s.Show("evens: ", evens)

	// newNums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	newNums := []int{0, 1, 2, 3, 4}
	// mid := len(newNums) / 2

	// midNums := newNums[1:3]
	midNums := append(newNums[1:3], 3, 4)
	// midNums := newNums[mid-1 : mid+1]

	s.Show("newNums: ", newNums)
	s.Show("midNums: ", midNums)

}
