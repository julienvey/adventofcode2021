package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func readInput(in string) string {
	file, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(file))
}
