package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/metdinov/aoc2020/utils"
)

type passport = map[string]string

func main() {
	f, err := os.Open("../data/day4a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(utils.SplitBlankFunc)
	validPassports := 0
	re := regexp.MustCompile(`(\w+):(#?\w+)`)

	for s.Scan() {
		line := s.Text()
		p := make(passport, 9)
		match := re.FindAllStringSubmatch(line, -1)

		for i := range match {
			p[match[i][1]] = match[i][2]
		}
		if isValidPassword(p) {
			validPassports++
		}
	}

	fmt.Println(validPassports)
}

type keyValidator = func(string) bool

var keyValidators = map[string]keyValidator{
	"byr": func(value string) bool { return validateYear(value, 1920, 2002) },
	"iyr": func(value string) bool { return validateYear(value, 2010, 2020) },
	"eyr": func(value string) bool { return validateYear(value, 2020, 2030) },
	"hgt": validateHeight,
	"hcl": validateHairColor,
	"ecl": validateEyeColor,
	"pid": validatePID,
}

func isValidPassword(p passport) bool {
	for key, validator := range keyValidators {
		if !validator(p[key]) {
			return false
		}
	}
	return true
}

func validateYear(str string, min, max int) bool {
	re := regexp.MustCompile(`^\d{4}$`)
	if re.MatchString(str) {
		number, err := strconv.Atoi(str)
		return err == nil && min <= number && number <= max
	}
	return false
}

func validateHeight(str string) bool {
	re := regexp.MustCompile(`^(\d+)(cm|in)$`)
	match := re.FindStringSubmatch(str)
	if match == nil {
		return false
	}

	height, err := strconv.Atoi(match[1])
	if err != nil {
		return false
	}

	if match[2] == "cm" {
		return 150 <= height && height <= 193
	}
	return 59 <= height && height <= 76
}

func validateHairColor(str string) bool {
	re := regexp.MustCompile(`^#[\d|a-f]{6}$`)
	return re.MatchString(str)
}

func validateEyeColor(str string) bool {
	switch str {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}

func validatePID(str string) bool {
	re := regexp.MustCompile(`^\d{9}$`)
	return re.MatchString(str)
}
