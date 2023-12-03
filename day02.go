package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day02(input []string) {
	part1 := 0
	part2 := 0
	for i, line := range input {
		samples := parseGame(line)
		if canUseLimitedNumberOfCubes(samples) {
			part1 += i +1
		}
		part2 += minPossibleCubesProduct(samples)
	}
	fmt.Println("Result part 1:", part1)
	fmt.Println("Result part 2:", part2)
}

func canUseLimitedNumberOfCubes(samples []map[string]int) bool {
	for _, sample := range samples {
		if sample["red"] > 12 || sample["green"] > 13 || sample["blue"] > 14 {
			return false
		}
	}
	return true
}

func minPossibleCubesProduct(samples []map[string]int) int {
	maxUsed := map[string]int{}
	for _, sample := range samples {
		for k, v := range sample {
			maxUsed[k] = max(maxUsed[k], v)
		}
	}
	res := 1
	for _, v := range maxUsed {
		res *= v
	}
	return res
}

func parseGame(raw string) []map[string]int {
	samples := []map[string]int{}
	for _, sample := range strings.Split(strings.SplitN(raw, ": ", 2)[1], "; ") {
		drawnCubes := map[string]int{}
		for _, cubes := range strings.Split(sample, ", ") {
			cube := strings.Split(cubes, " ")
			nr, _ := strconv.Atoi(cube[0])
			drawnCubes[cube[1]] = nr
		}
		samples = append(samples, drawnCubes)
	}
	return samples
}
