package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readInput(in string) [][]int {
	file, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}
	var input [][]int
	for _, line := range strings.Split(string(file), "\n") {
		if line == "" {
			continue
		}
		var inputLine []int
		for _, c := range strings.TrimSpace(line) {
			conv, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}
			inputLine = append(inputLine, conv)
		}
		input = append(input, inputLine)
	}
	return input
}
