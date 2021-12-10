package utils

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadInput(in string) []int {
	file, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}
	var input []int
	for _, line := range strings.Split(string(file), ",") {
		if line == "" {
			continue
		}
		inputInt, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, inputInt)
	}
	return input
}

func ReadInputString(in string) []string {
	file, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}
	var input []string
	for _, line := range strings.Split(string(file), "\n") {
		if line == "" {
			continue
		}
		input = append(input, line)
	}
	return input
}
