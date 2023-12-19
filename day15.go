package main

import (
	"fmt"
	"strings"
)

type lens struct {
	label string
	focal int
}

func Day15(input []string) {
	part1 := 0
	for _, s := range strings.Split(input[0], ",") {
		part1 += hash(s)
	}
	boxes := make([][]lens, 256)
	for _, s := range strings.Split(input[0], ",") {
		if s[len(s)-1] == '-' {
			label := s[:len(s)-1]
			h := hash(label)
			if ind := index(boxes[h], label); ind != -1 {
				boxes[h] = Remove(boxes[h], ind)
			}
		} else {
			label := s[:len(s)-2]
			h := hash(label)
			lab := lens{label, int(s[len(s)-1] - '0')}
			if ind := index(boxes[h], label); ind != -1 {
				boxes[h][ind] = lab
			} else {
				boxes[h] = append(boxes[h], lab)
			}
		}
	}

	part2 := 0
	for bi, box := range boxes {
		for slot, l := range box {
			part2 += (bi+1)*(slot+1)*l.focal
		}
	}
	fmt.Println("Result part 1:", part1)
	fmt.Println("Result part 2:", part2)
}

func hash(s string) int {
	h := 0
	for _, c := range s {
		h = (h + int(c)) * 17 % 256
	}
	return h
}

func index(box []lens, label string) int {
	for i, l := range box {
		if l.label == label {
			return i
		}
	}
	return -1
}
