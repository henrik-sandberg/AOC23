package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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
	digits := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	re := regexp.MustCompile("\\d|" + strings.Join(digits[1:], "|"))
	tr := func(s string) string {
		if i := IndexOf(digits, s); i != -1 {
			return strconv.Itoa(i)
		}
		return s
	}
	for _, line := range inp {
		for _, digit := range digits {
			line = strings.ReplaceAll(line, digit, digit+digit[len(digit)-1:])
		}
		all := re.FindAllString(line, -1)
		n, _ := strconv.Atoi(tr(all[0]) + tr(all[len(all)-1]))
		res += n
	}
	return
}
