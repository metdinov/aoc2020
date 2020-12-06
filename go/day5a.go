package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/metdinov/aoc2020/utils"
)

func main() {
	f, err := os.Open("../data/day5a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	maxSeatID := 0

	for s.Scan() {
		seat := s.Text()
		row, col := getRowAndColumn(seat)

		seatID := row*8 + col
		maxSeatID = utils.Max(seatID, maxSeatID)
	}

	fmt.Println(maxSeatID)
}

func getRowAndColumn(seat string) (int, int) {
	row, col := 0, 0
	for i := 6; i >= 0; i-- {
		if seat[6-i] == 'B' {
			row += 1 << i
		}
	}
	for j := 2; j >= 0; j-- {
		if seat[9-j] == 'R' {
			col += 1 << j
		}
	}
	return row, col
}
