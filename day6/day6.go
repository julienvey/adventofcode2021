package main

import (
	"fmt"
	"github.com/julienvey/adventofcode2021/utils"
)

func main() {
	input := utils.ReadInput("day6/input.txt")
	fmt.Printf("Solution Day 6, Part 1: %d\n", getGrowth(input, 80))
	fmt.Printf("Solution Day 6, Part 2: %d\n", getGrowth(input, 256))
}

func getGrowth(input []int, days int) int {
	count := map[int]int{}
	for _, i := range input {
		count[i]++
	}

	for d := 0; d < days; d++ {
		grow := count[0]
		count[0] = count[1]
		count[1] = count[2]
		count[2] = count[3]
		count[3] = count[4]
		count[4] = count[5]
		count[5] = count[6]
		count[6] = count[7] + grow
		count[7] = count[8]
		count[8] = grow
	}
	var sum int
	for _, v := range count {
		sum += v
	}
	return sum
}
