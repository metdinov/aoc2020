package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/metdinov/aoc2020/utils"
)

func main() {
	f, err := os.Open("../data/day1b.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	seen := utils.NewIntSet()

	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}

		if num, found := seen.Find(containsComplement(seen, i)); found {
			fmt.Println(num * (2020 - i - num) * i)
			return
		}
		seen.Add(i)
	}
}

func containsComplement(set *utils.IntSet, num int) func(int) bool {
	return func(el int) bool {
		return (el != 2020-num) && set.Contains(2020-num-el)
	}
}
