package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../data/day2b.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	valid := 0
	var pos1, pos2 int
	var char byte
	var passwd string

	for s.Scan() {
		fmt.Sscanf(s.Text(), "%d-%d %c: %s", &pos1, &pos2, &char, &passwd)
		if (passwd[pos1-1] == char) != (passwd[pos2-1] == char) {
			valid++
		}
	}
	fmt.Println(valid)
}
