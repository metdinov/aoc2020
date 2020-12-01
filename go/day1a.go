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
	f, err := os.Open("../data/day1a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	complements := utils.NewIntSet()

	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}

		if complements.Contains(i) {
			fmt.Println(i * (2020 - i))
			return
		}
		complements.Add(2020 - i)
	}
}
