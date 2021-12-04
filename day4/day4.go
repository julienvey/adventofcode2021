package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := readInput("day4/input.txt")
	fmt.Printf("Solution Day 4, Part 1: %d\n", SolvePuzzlePart1(input))
	fmt.Printf("Solution Day 4, Part 2: %d\n", SolvePuzzlePart2(input))
}

type Input struct {
	Numbers []int
	Boards  [][][]BoardNumber
}

type BoardNumber struct {
	Number int
	Marked bool
}

func readInput(in string) Input {
	file, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}
	var input Input
	split := strings.Split(string(file), "\n\n")

	for _, num := range strings.Split(split[0], ",") {
		numInt, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		input.Numbers = append(input.Numbers, numInt)
	}

	for _, boardRaw := range split[1:] {
		board := make([][]BoardNumber, 5)
		for i := range board {
			board[i] = make([]BoardNumber, 5)
		}
		for rowNum, row := range strings.Split(boardRaw, "\n") {
			for colNum, num := range strings.Fields(row) {
				if num == "" {
					continue
				}
				numInt, err := strconv.Atoi(num)
				if err != nil {
					log.Fatal(err)
				}
				board[rowNum][colNum] = BoardNumber{
					Number: numInt,
				}
			}
		}
		input.Boards = append(input.Boards, board)
	}
	return input
}

func SolvePuzzlePart1(input Input) int {
	for _, draw := range input.Numbers {
		markBoards(input.Boards, draw)
		if board, _ := checkWinner(input.Boards); board != nil {
			return getScore(board, draw)
		}

	}
	log.Fatal("could not find winner")
	return 0
}

func SolvePuzzlePart2(input Input) int {
	for _, draw := range input.Numbers {
		markBoards(input.Boards, draw)
		board, index := checkWinner(input.Boards)
		if len(input.Boards) == 1 && board != nil {
			return getScore(input.Boards[0], draw)
		}
		for board != nil {
			input.Boards = append(input.Boards[:index], input.Boards[index+1:]...)
			board, index = checkWinner(input.Boards)
		}

	}
	log.Fatal("some board never wins")
	return 0
}

func getScore(board [][]BoardNumber, draw int) int {
	var sum int
	for _, row := range board {
		for _, item := range row {
			if !item.Marked {
				sum += item.Number
			}
		}
	}
	return sum * draw
}

func checkWinner(boards [][][]BoardNumber) ([][]BoardNumber, int) {
	for boardIndex, board := range boards {
		for i := 0; i < 5; i++ {
			allRowsMarked := true
			allColMarked := true
			for j := 0; j < 5; j++ {
				allRowsMarked = allRowsMarked && board[i][j].Marked
				allColMarked = allColMarked && board[j][i].Marked
			}
			if allColMarked || allRowsMarked {
				return board, boardIndex
			}
		}
	}
	return nil, 0
}

func markBoards(boards [][][]BoardNumber, draw int) {
	for i, _ := range boards {
		for j, _ := range boards[i] {
			for k, _ := range boards[i][j] {
				if boards[i][j][k].Number == draw {
					boards[i][j][k].Marked = true
				}
			}
		}
	}
}
