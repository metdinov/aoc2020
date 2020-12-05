package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/metdinov/aoc2020/utils"
)

func main() {
	f, err := os.Open("../data/day4a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(utils.SplitBlankFunc)
	validPassports := 0
	re := regexp.MustCompile(`(\w+):`)

	for s.Scan() {
		line := s.Text()
		seenKeys := utils.NewStringSet()
		match := re.FindAllStringSubmatch(line, -1)

		for i := range match {
			seenKeys.Add(match[i][1])
		}
		if hasRequiredKeys(seenKeys) {
			validPassports++
		}
	}

	fmt.Println(validPassports)
}

func hasRequiredKeys(seenKeys *utils.StringSet) bool {
	return seenKeys.ContainsAll("byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid")
}
