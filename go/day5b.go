package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("../data/day5a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	seatIDs := make([]int, 128*8)
	i := 0

	for s.Scan() {
		seat := s.Text()
		row, col := getRowAndColumn(seat)
		seatIDs[i] = row*8 + col
		i++
	}
	fmt.Println(findMySeat(seatIDs))
}

func getRowAndColumn(seat string) (int, int) {
	row, col := 0, 0
	for i := 6; i >= 0; i-- {
		if seat[6-i] == 'B' {
			row += (1 << i)
		}
	}
	for j := 2; j >= 0; j-- {
		if seat[9-j] == 'R' {
			col += (1 << j)
		}
	}
	return row, col
}

func findMySeat(seatIDs []int) int {
	sort.Ints(seatIDs)
	seatsCount := len(seatIDs)
	for i := len(seatIDs) / 2; i < seatsCount; i++ {
		if seatIDs[i+1]-seatIDs[i] == 2 {
			return seatIDs[i] + 1
		}
		if seatIDs[seatsCount-i]-seatIDs[seatsCount-i-1] == 2 {
			return seatIDs[seatsCount-i] + 1
		}
	}
	return -1
}
