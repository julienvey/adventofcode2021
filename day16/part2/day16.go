package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	input := readInput("day16/input.txt")
	fmt.Printf("Solution Day 16, Part 2: %d\n", SolvePuzzlePart2(input))
}

func SolvePuzzlePart2(input string) int64 {
	bin := hex2Bin(input)

	res, _ := handleGenericPacket(bin)
	return res
}

func handleGenericPacket(bin string) (int64, string) {
	bin = bin[3:]
	packetType := toInt(bin[0:3])
	bin = bin[3:]
	var values []int64
	switch packetType {
	case 0: // Sum
		values, bin = handleOperator(bin)
		sum := values[0]
		for _, v := range values[1:] {
			sum += v
		}
		return sum, bin
	case 1: // Product
		values, bin = handleOperator(bin)
		product := values[0]
		for _, v := range values[1:] {
			product = product * v
		}
		return product, bin
	case 2: // Minimum
		values, bin = handleOperator(bin)
		min := values[0]
		for _, v := range values[1:] {
			if v < min {
				min = v
			}
		}
		return min, bin
	case 3: // Maximum
		values, bin = handleOperator(bin)
		max := values[0]
		for _, v := range values[1:] {
			if v > max {
				max = v
			}
		}
		return max, bin
	case 4: // Single Digit
		return handleSinglePacket(bin)
	case 5: // Greater Than
		values, bin = handleOperator(bin)
		var num int64
		if values[0] > values[1] {
			num = 1
		}
		return num, bin
	case 6: // Less Than
		values, bin = handleOperator(bin)
		var num int64
		if values[0] < values[1] {
			num = 1
		}
		return num, bin
	case 7: // Equal
		values, bin = handleOperator(bin)
		var num int64
		if values[0] == values[1] {
			num = 1
		}
		return num, bin
	default:
		log.Fatal("Should not happen")
		return 0, ""
	}
}

func handleOperator(bin string) ([]int64, string) {
	var values []int64
	lengthTypeID := toInt(string(bin[0]))
	bin = bin[1:]
	if lengthTypeID == 0 {
		lenPackets := toInt(bin[0:15])
		bin = bin[15:]
		subpackets := bin[:lenPackets]
		for !containsOnlyZero(subpackets) {
			var value int64
			value, subpackets = handleGenericPacket(subpackets)
			values = append(values, value)
		}
		bin = bin[lenPackets:]
	} else {
		numberOfPackets := toInt(bin[0:11])
		bin = bin[11:]
		for i := 0; i < int(numberOfPackets); i++ {
			var value int64
			value, bin = handleGenericPacket(bin)
			values = append(values, value)
		}
	}
	return values, bin
}

func handleSinglePacket(bin string) (int64, string) {
	var digit string
	for {
		packet := bin[0:5]
		digit += packet[1:]
		bin = bin[5:]
		if packet[0] == '0' {
			break
		}
	}
	return toInt(digit), bin
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

func toInt(s string) int64 {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
