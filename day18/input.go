package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Pair struct {
	parent  *Pair
	left    *Pair
	right   *Pair
	literal int
}

func readInput(in string) []*Pair {
	content, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(content), "\n")

	pairs := createPairs(input)
	return pairs
}

func createPairs(input []string) []*Pair {
	var pairs []*Pair
	for _, line := range input {
		if line == "" {
			continue
		}
		pair, rest := ParsePair(line)
		if len(rest) != 0 {
			panic(line)
		}
		pairs = append(pairs, pair)
	}
	return pairs
}

var NumPat = regexp.MustCompile("\\d+")

func ParsePair(text string) (*Pair, string) {
	if text[0] == '[' {
		p1, text := ParsePair(text[1:])
		if text[0] != ',' {
			panic(text)
		}
		p2, text := ParsePair(text[1:])
		if text[0] != ']' {
			panic(text)
		}
		parent := Pair{left: p1, right: p2}
		p1.parent = &parent
		p2.parent = &parent
		return &parent, text[1:]
	}
	pos := NumPat.FindStringIndex(text)
	if pos == nil {
		panic(text)
	}
	value, err := strconv.Atoi(text[pos[0]:pos[1]])
	if err != nil {
		panic(text)
	}
	return &Pair{literal: value}, text[pos[1]:]
}
