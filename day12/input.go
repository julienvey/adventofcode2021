package main

import (
	"io/ioutil"
	"log"
	"strings"
)

type Entry struct {
	Source string
	Dest   string
}

func readInput(in string) []Entry {
	file, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}
	var input []Entry
	for _, line := range strings.Split(string(file), "\n") {
		if line == "" {
			continue
		}
		entrySplit := strings.Split(line, "-")
		input = append(input, Entry{
			Source: entrySplit[0],
			Dest:   entrySplit[1],
		})

	}
	return input
}
