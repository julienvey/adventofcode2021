package main

import (
	"io/ioutil"
	"log"
	"strings"
)

type Input struct {
	Start string
	Map   map[string]string
}

func readInput(in string) Input {
	file, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}
	input := Input{
		Map: make(map[string]string),
	}
	for _, line := range strings.Split(string(file), "\n") {
		if line == "" {
			continue
		}

		if strings.Contains(line, " -> ") {
			split := strings.Split(line, " -> ")
			input.Map[split[0]] = split[1]
		} else {
			input.Start = line
		}
	}
	return input
}
