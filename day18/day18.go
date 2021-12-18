package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := readInput("day18/input.txt")
	fmt.Printf("Solution Day 18, Part 1: %d\n", GetMagnitude(SolveAddition(input)))
	fmt.Printf("Solution Day 18, Part 2: %d\n", LargestMagnitude("day18/input.txt"))
}

func LargestMagnitude(in string) int {
	l := len(readInput(in))
	var largestMag int
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i == j {
				continue
			}
			input := readInput(in)
			mag := GetMagnitude(SolveAddition([]*Pair{input[i], input[j]}))
			if mag > largestMag {
				largestMag = mag
			}
		}
	}
	return largestMag
}

func SolveAddition(input []*Pair) *Pair {
	current := input[0]
	for i := 1; i < len(input); i++ {
		current = Sum(current, input[i])
		Reduce(current)
	}
	return current
}

func GetMagnitude(pair *Pair) int {
	if pair.left == nil {
		return pair.literal
	}

	return GetMagnitude(pair.left)*3 + GetMagnitude(pair.right)*2
}

func Sum(a *Pair, b *Pair) *Pair {
	sum := &Pair{
		left:  a,
		right: b,
	}
	a.parent = sum
	b.parent = sum
	return sum
}

func (p Pair) String() string {
	if p.left != nil {
		return fmt.Sprintf("[%s,%s]", p.left.String(), p.right.String())
	}
	return strconv.Itoa(p.literal)
}

func findClosestLeftPair(pair *Pair) *Pair {
	currentPair := pair
	for {
		if currentPair.parent == nil {
			return nil
		}
		if currentPair.parent.left == currentPair {
			currentPair = currentPair.parent
			continue
		}
		searchAncester := currentPair.parent.left
		for {
			if searchAncester.right == nil {
				return searchAncester
			}
			searchAncester = searchAncester.right
		}
	}
}

func findClosestRightPair(pair *Pair) *Pair {
	currentPair := pair
	for {
		if currentPair.parent == nil {
			return nil
		}
		if currentPair.parent.right == currentPair {
			currentPair = currentPair.parent
			continue
		}
		searchAncester := currentPair.parent.right
		for {
			if searchAncester.left == nil {
				return searchAncester
			}
			searchAncester = searchAncester.left
		}
	}
}

func Reduce(pair *Pair) *Pair {
	for {
		firstNested := findFirstNestedPair(pair, 0)
		if firstNested != nil {
			closestLeft := findClosestLeftPair(firstNested)
			if closestLeft != nil {
				if closestLeft.left != nil || closestLeft.right != nil {
					panic("Closest left should be nil")
				}
				closestLeft.literal += firstNested.left.literal
			}
			closestRight := findClosestRightPair(firstNested)
			if closestRight != nil {
				if closestRight.left != nil || closestRight.right != nil {
					panic("Closest left should be nil")
				}
				closestRight.literal += firstNested.right.literal
			}
			firstNested.left = nil
			firstNested.right = nil
			firstNested.literal = 0
			continue
		}
		firstOver10 := findFirstOverTen(pair)
		if firstOver10 != nil {
			firstOver10.left = &Pair{
				parent:  firstOver10,
				literal: firstOver10.literal / 2,
			}
			firstOver10.right = &Pair{
				parent:  firstOver10,
				literal: (firstOver10.literal + 1) / 2,
			}
			firstOver10.literal = 0
			continue
		}
		return pair
	}
}

func findFirstOverTen(pair *Pair) *Pair {
	if pair.left == nil {
		if pair.literal >= 10 {
			return pair
		}
		return nil
	}

	if nestedPair := findFirstOverTen(pair.left); nestedPair != nil {
		return nestedPair
	}

	if nestedPair := findFirstOverTen(pair.right); nestedPair != nil {
		return nestedPair
	}

	return nil

}

func findFirstNestedPair(pair *Pair, depth int) *Pair {
	if pair.left == nil {
		return nil
	}

	if pair.left.left == nil && depth == 4 {
		return pair
	}

	if nestedPair := findFirstNestedPair(pair.left, depth+1); nestedPair != nil {
		return nestedPair
	}

	if nestedPair := findFirstNestedPair(pair.right, depth+1); nestedPair != nil {
		return nestedPair
	}

	return nil
}
