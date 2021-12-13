package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Input struct {
	Paper        []map[string]int
	Instructions []Instruction
}

type Instruction struct {
	Direction string
	Count     int
}

func readInput(in string) Input {
	file, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}
	var input Input
	for _, line := range strings.Split(string(file), "\n") {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "fold along ") {
			line = strings.ReplaceAll(line, "fold along ", "")
			split := strings.Split(line, "=")
			val, err := strconv.Atoi(split[1])
			if err != nil {
				log.Fatal(err)
			}
			input.Instructions = append(input.Instructions, Instruction{Direction: split[0], Count: val})
		} else {
			split := strings.Split(line, ",")
			x, err := strconv.Atoi(split[0])
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.Atoi(split[1])
			if err != nil {
				log.Fatal(err)
			}
			input.Paper = append(input.Paper, map[string]int{"x": x, "y": y})
		}
	}
	return input
}
