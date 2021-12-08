package main

import (
	"io/ioutil"
	"log"
	"strings"
)

type Entry struct {
	Inputs  []string
	Outputs []string
}

func readInput(in string) []Entry {
	file, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}
	var entries []Entry
	for _, line := range strings.Split(string(file), "\n") {
		if line == "" {
			continue
		}
		item := strings.Split(line, "|")
		if len(item) != 2 {
			log.Fatal("entry should have two items, separated by |")
		}

		var entry Entry
		for _, signal := range strings.Split(strings.TrimSpace(item[0]), " ") {
			entry.Inputs = append(entry.Inputs, strings.TrimSpace(signal))
		}
		for _, digit := range strings.Split(strings.TrimSpace(item[1]), " ") {
			entry.Outputs = append(entry.Outputs, strings.TrimSpace(digit))
		}
		entries = append(entries, entry)
	}
	return entries
}
