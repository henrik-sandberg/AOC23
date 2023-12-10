package main

import (
	"fmt"
	"strings"
)

type tuple struct {
	left, right string
}

func Day08(input []string) {
	dirs := input[0]
	network := parseNetwork(input[2:])
	fmt.Println("Result part 1:", day08Part1(dirs, network))
	fmt.Println("Result part 2:", day08Part2(dirs, network))
}

func day08Part1(dirs string, network map[string]tuple) int {
	return stepsUntilTarget("AAA", "ZZZ", dirs, network)
}

func day08Part2(dirs string, network map[string]tuple) int {
	paths := []int{}
	for k, _ := range network {
		if strings.HasSuffix(k, "A") {
			paths = append(paths, stepsUntilTarget(k, "Z", dirs, network))
		}
	}
	return lcm(paths...)
}

func stepsUntilTarget(start, targetSuffix, dirs string, network map[string]tuple) (res int) {
	next := start
	for !strings.HasSuffix(next, targetSuffix) {
		dir := string(dirs[res % len(dirs)])
		if dir == "L" {
			next = network[next].left
		} else {
			next = network[next].right
		}
		res++
	}
	return
}

func parseNetwork(raw []string) map[string]tuple {
	network := make(map[string]tuple, len(raw))
	for _, line := range raw {
		split := strings.Split(line, " = ")
		vals := split[1][1:len(split[1])-1]
		lr := strings.Split(vals, ", ")
		network[split[0]] = tuple{lr[0], lr[1]}
	}
	return network
}
