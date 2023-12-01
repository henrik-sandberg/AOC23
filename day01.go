package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var digits = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func Day01(input []string) {
	fmt.Println("Result part 1:", day01_part1(input))
	fmt.Println("Result part 2:", day01_part2(input))
}

func day01_part1(inp []string) (res int) {
	re := regexp.MustCompile("\\d")
	for _, line := range inp {
		all := re.FindAllString(line, -1)
		n, _ := strconv.Atoi(all[0] + all[len(all)-1])
		res += n
	}
	return
}

func day01_part2(inp []string) (res int) {
	re := regexp.MustCompile("\\d|" + strings.Join(digits[1:], "|"))
	for _, line := range inp {
		for _, digit := range digits {
			line = strings.ReplaceAll(line, digit, digit+string(digit[len(digit)-1]))
		}
		all := re.FindAllString(line, -1)
		first := all[0]
		if i := IndexOf(digits, first); i != -1 {
			first = strconv.Itoa(i)
		}
		last := all[len(all)-1]
		if i := IndexOf(digits, last); i != -1 {
			last = strconv.Itoa(i)
		}
		n, _ := strconv.Atoi(first + last)
		res += n
	}
	return
}
