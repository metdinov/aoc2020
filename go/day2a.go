package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("../data/day2a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	valid := 0
	var min, max int
	var char byte
	var passwd string

	for s.Scan() {
		fmt.Sscanf(s.Text(), "%d-%d %c: %s", &min, &max, &char, &passwd)
		count := strings.Count(passwd, string(char))

		if count >= min && count <= max {
			valid++
		}
	}

	fmt.Println(valid)
}
