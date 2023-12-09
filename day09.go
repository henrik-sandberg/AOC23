package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day09(input []string) {
	part1 := 0
	part2 := 0
	for _, line := range input {
		numbers := strings.Fields(line)
		ns := make([]int, len(numbers))
		for i := range numbers {
			n, _ := strconv.Atoi(numbers[i])
			ns[i] = n
		}
		a, b := nextPolynomial(ns)
		part1 += b
		part2 += a
	}
	fmt.Println("Result part 1", part1)
	fmt.Println("Result part 2", part2)
}

func nextPolynomial(numbers []int) (int, int) {
	first := numbers[0]
	last := numbers[len(numbers)-1]
	if constantDiff(numbers) {
		return first, last
	}
	diffs := make([]int, len(numbers)-1)
	for i := range numbers[:len(numbers)-1] {
		diffs[i] = numbers[i+1]-numbers[i]
	}
	a, b := nextPolynomial(diffs)
	return first - a, last + b
}

func constantDiff(numbers []int) bool {
	for i := range numbers[:len(numbers)-2] {
		if numbers[i+2] - numbers[i+1] != numbers[i+1] - numbers[0] {
			return false
		}
	}
	return true
}
