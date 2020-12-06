package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/metdinov/aoc2020/utils"
)

func main() {
	f, err := os.Open("../data/day6a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(utils.SplitBlankFunc)
	total := 0

	for s.Scan() {
		line := s.Text()
		questions := utils.NewByteSet()
		for i := 0; i < len(line); i++ {
			if line[i] == '\n' {
				continue
			}
			questions.Add(line[i])
		}
		total += questions.Size()
	}

	fmt.Println(total)
}
