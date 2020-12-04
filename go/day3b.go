package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type slope struct {
	right      int
	down       int
	currentPos int
	treeCount  int
}

func main() {
	f, err := os.Open("../data/day3a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	slopes := []*slope{
		&slope{right: 1, down: 1},
		&slope{right: 3, down: 1},
		&slope{right: 5, down: 1},
		&slope{right: 7, down: 1},
		&slope{right: 1, down: 2},
	}

	lineLength := 0
	lineNumber := 0

	for s.Scan() {
		line := s.Text()
		if lineLength == 0 {
			lineLength = len(line)
		}

		for _, sl := range slopes {
			if isTreeA(line, lineNumber, sl) {
				sl.treeCount = sl.treeCount + 1
			}
			if lineNumber%sl.down == 0 {
				sl.currentPos = (sl.currentPos + sl.right) % lineLength
			}
		}
		lineNumber++
	}

	answer := 1
	for _, sl := range slopes {
		answer *= sl.treeCount
	}
	fmt.Println(answer)
}

func isTreeA(line string, lineNumber int, sl *slope) bool {
	return (lineNumber%sl.down == 0) && (line[sl.currentPos] == '#')
}
