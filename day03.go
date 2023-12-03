package main

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

func Day03(input []string) {
	grid := asMap(input)
	numbers := findAllNumbers(input)
	part1 := 0
	part2 := 0
	for point, char := range grid {
		if char != '.' && !unicode.IsDigit(char) {
			sni := getSurroundingNumberIndexes(point, grid, numbers)
			sn := make([]int, 0, len(sni))
			for ni, _ := range sni {
				sn = append(sn, createNumber(ni, input))
			}
			part1 += sum(sn...)
			if len(sn) == 2 {
				part2 += multiply(sn...)
			}
		}
	}
	fmt.Println("Result part 1:", part1)
	fmt.Println("Result part 2:", part2)
}

func asMap(input []string) map[point2d]rune {
	m := map[point2d]rune{}
	for y, row := range input {
		for x, col := range row {
			m[point2d{x, y}] = col
		}
	}
	return m
}

func findAllNumbers(input []string) [][]rectangle {
	re := regexp.MustCompile("\\d+")
	ret := make([][]rectangle, len(input))
	for ri, row := range input {
		for _, match := range re.FindAllStringIndex(row, -1) {
			ret[ri] = append(ret[ri], rectangle {
				point2d {match[0], ri},
				point2d {match[1], ri},
			})
		}
	}
	return ret
}

func getSurroundingNumberIndexes(point point2d, grid map[point2d]rune, numbers [][]rectangle) map[rectangle]bool {
	ret := map[rectangle]bool{}
	for p := range point.adjacent() {
		if unicode.IsDigit(grid[p]) {
			for _, numberIndexes := range numbers[p.y] {
				start := numberIndexes.left
				end := numberIndexes.right
				if start.x <= p.x && p.x <= end.x {
					ret[rectangle {left: start, right: end}] = true
				}
			}
		}
	}
	return ret
}

func createNumber(numberIndexes rectangle, input []string) int {
	start := numberIndexes.left
	end := numberIndexes.right
	if start.y != end.y {
		panic("A number can only be on a single line")
	}
	number, _ := strconv.Atoi(input[start.y][start.x:end.x])
	return number
}
