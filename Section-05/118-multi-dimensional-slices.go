package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	spendings := Fetch()

	fmt.Println(spendings[1])

	var spendTot int
	for i, line := range spendings {
		daySpend := addSpending(line)
		fmt.Printf("Total spend on day #%d:\t%d\n", i+1, daySpend)
		spendTot += daySpend
	}
	fmt.Printf("Total spending:\t\t%d\n", spendTot)

}

func addSpending(spends []int) int {
	var total int

	for _, spend := range spends {
		total += spend
	}

	return total
}

func Fetch() [][]int {
	content := `200 100
25 10 45 60
5 15 35
95 10
50 25`

	lines := strings.Split(content, "\n")
	ret := make([][]int, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)
		ret[i] = make([]int, len(fields))

		for j, field := range fields {
			// ret[i][j], _ = strconv.Atoi(field)
			spending, err := strconv.Atoi(field)
			if err != nil {
				fmt.Printf("error converting to int: %v\n", err)
			}
			ret[i][j] = spending

		}
	}

	return ret
}
