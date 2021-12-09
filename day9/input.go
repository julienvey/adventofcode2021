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
	var cave [][]int
	for _, line := range strings.Split(string(file), "\n") {
		if line == "" {
			continue
		}
		var caveLine []int
		for _, c := range strings.TrimSpace(line) {
			conv, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}
			caveLine = append(caveLine, conv)
		}
		cave = append(cave, caveLine)
	}
	return cave
}
