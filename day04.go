package main

import (
	"fmt"
	"strings"
)

func Day04(input []string) {
	part1 := 0
	part2 := make([]int, len(input))
	for i := range part2 {
		part2[i] = 1
	}
	for i, row := range input {
		cards := strings.Split(strings.SplitN(row, ": ", 2)[1], " | ")
		wins := len(Intersect(strings.Fields(cards[0]), strings.Fields(cards[1])))
		if wins > 0 {
			part1 += 1 << (wins - 1)
		}
		for j := i + 1; j < i+wins+1; j++ {
			part2[j] += part2[i]
		}
	}
	fmt.Println(part1)
	fmt.Println(sum(part2...))
}
