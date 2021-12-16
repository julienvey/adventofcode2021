package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	input := readInput("day16/input.txt")
	fmt.Printf("Solution Day 16, Part 1: %d\n", SolvePuzzlePart1(input))
}

func SolvePuzzlePart2(input string) int {
	return 0
}

func SolvePuzzlePart1(input string) int {
	bin := hex2Bin(input)

	sum, _ := handleGenericPacket(bin)
	return sum
}

func handleGenericPacket(bin string) (int, string) {
	packetVersion := toInt(bin[0:3])
	versionSum := packetVersion
	bin = bin[3:]
	packetType := toInt(bin[0:3])
	bin = bin[3:]
	switch packetType {
	case 4:
		bin = handleSinglePacket(bin)
		return versionSum, bin
	default:
		var count int
		count, bin = handleOperator(bin)
		return count + versionSum, bin
	}
}

func handleOperator(bin string) (int, string) {
	var sum int
	lengthTypeID := toInt(string(bin[0]))
	bin = bin[1:]
	if lengthTypeID == 0 {
		lenPackets := toInt(bin[0:15])
		bin = bin[15:]
		subpackets := bin[:lenPackets]
		for !containsOnlyZero(subpackets) {
			var count int
			count, subpackets = handleGenericPacket(subpackets)
			sum += count
		}
		bin = bin[lenPackets:]
	} else {
		numberOfPackets := toInt(bin[0:11])
		bin = bin[11:]
		for i := 0; i < numberOfPackets; i++ {
			var count int
			count, bin = handleGenericPacket(bin)
			sum += count
		}
	}
	return sum, bin
}

func handleSinglePacket(bin string) string {
	for {
		packet := bin[0:5]
		bin = bin[5:]
		if packet[0] == '0' {
			break
		}
	}
	return bin
}

func containsOnlyZero(subpackets string) bool {
	for _, c := range subpackets {
		if c != '0' {
			return false
		}
	}
	return true
}

func hex2Bin(h string) string {
	res := ""
	for _, n := range h {
		b, err := strconv.ParseInt(string(n), 16, 64)
		if err != nil {
			log.Fatal(err)
		}
		res += fmt.Sprintf("%04b", b)
	}
	return res
}

func toInt(s string) int {
	i, err := strconv.ParseInt(s, 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	return int(i)
}
