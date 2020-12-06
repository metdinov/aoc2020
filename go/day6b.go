package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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
		line := strings.Trim(s.Text(), "\n")
		questions := make(map[byte]int)
		persons := 1
		for i := 0; i < len(line); i++ {
			if line[i] == '\n' {
				persons++
				continue
			}
			questions[line[i]]++
		}

		totalKeys := 0
		for _, count := range questions {
			if count == persons {
				total++
				totalKeys++
			}
		}
	}

	fmt.Println(total)
}
