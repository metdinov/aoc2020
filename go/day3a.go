package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../data/day3a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	currentPos := 0
	treeCount := 0
	lineLength := 0

	for s.Scan() {
		line := s.Text()
		if lineLength == 0 {
			lineLength = len(line)
		}

		if isTree(line[currentPos]) {
			treeCount++
		}

		currentPos = (currentPos + 3) % lineLength
	}

	fmt.Println(treeCount)
}

func isTree(char byte) bool {
	return char == '#'
}
