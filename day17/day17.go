package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	part1 := SolvePuzzle("target area: x=81..129, y=-150..-108")
	fmt.Printf("Solution Day 17, Part 1: %d\n", part1)
}

type Target struct {
	xMin, xMax, yMin, yMax int
}

func SolvePuzzle(input string) int {
	fmt.Println(input)
	rex := regexp.MustCompile(`^target area: x=([-0-9]+)\.\.([-0-9]+), y=([-0-9]+)\.\.([-0-9]+)$`)
	res := rex.FindAllStringSubmatch(input, -1)[0]

	target := Target{}
	target.xMin, _ = strconv.Atoi(res[1])
	target.xMax, _ = strconv.Atoi(res[2])
	target.yMin, _ = strconv.Atoi(res[3])
	target.yMax, _ = strconv.Atoi(res[4])

	possiblesX := getXInterval(target)
	fmt.Println(possiblesX)

	// Je teste tous les X possible
	var globalMaxHeight int
	for _, possibleX := range possiblesX {
		fmt.Printf("Testing x: %d\n", possibleX)
		possibleY := target.yMin - 1
		// Je teste tous les Y possible, en démarrant par yMin, jusqu'à 1000 ¯\_(ツ)_/¯
		for {
			var xPos, yPos, maxHeight int
			var yFailed, tooHigh bool
			possibleY++
			yVeloc := possibleY
			xVeloc := possibleX

			// Je lance
			for {
				xPos += xVeloc
				yPos += yVeloc
				if inTarget(target, xPos, yPos) {
					if yPos > maxHeight {
						maxHeight = yPos
					}
					break
				}
				if missedTarget(target, xPos, yPos) {
					yFailed = true
					break
				}
				if yPos > maxHeight {
					maxHeight = yPos
				}
				if xVeloc > 0 {
					xVeloc--
				}
				yVeloc--
			}
			if !yFailed && !tooHigh && (maxHeight > globalMaxHeight) {
				globalMaxHeight = maxHeight
			}
			if possibleY > 1000 {
				break
			}
		}
	}
	return globalMaxHeight
}

func inTarget(target Target, x, y int) bool {
	return x >= target.xMin && x <= target.xMax && y >= target.yMin && y <= target.yMax
}

func missedTarget(target Target, x, y int) bool {
	return y < target.yMin || x > target.xMax
}

func getXInterval(target Target) []int {
	var possibles []int
	for x := 1; x < target.xMin; x++ {
		pos := x
		xval := x
		for {
			if pos >= target.xMin && pos <= target.xMax {
				possibles = append(possibles, x)
				break
			}
			if pos > target.xMax || xval == 0 {
				break
			}
			xval--
			pos += xval

		}
	}
	for x := target.xMin; x <= target.xMax; x++ {
		if !contains(possibles, x) {
			possibles = append(possibles, x)
		}

	}
	return possibles
}

func contains(arr []int, x int) bool {
	for _, i := range arr {
		if i == x {
			return true
		}
	}
	return false
}
