package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func readInput(in string) [][]int {
	content, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(content), "\n")
	input = input[:len(input)-1]

	grid := make([][]int, 0)
	for _, i := range input {
		row := make([]int, 0)
		for _, v := range i {
			c := v - '0'
			row = append(row, int(c))
		}
		grid = append(grid, row)
	}
	return grid
}
