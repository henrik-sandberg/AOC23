package main

import(
	"fmt"
	"strings"
)

type workflow struct {
	evaluate func(map[string]int) (string, bool)
}

func Day19(input []string) {
	workflows := make(map[string][]workflow)
	index := 0
	for ; input[index] != ""; index++ {
		s := input[index]
		ind := strings.Index(s, "{")
		workflows[s[:ind]] = parseWorkflows(s[ind:])
	}
	index++
	part1 := 0
	for ; index < len(input); index++ {
		params := parseParams(input[index])
		part1 += day19_part1(workflows, params)
	}
	fmt.Println("Result part 1", part1)
}

func day19_part1(wf map[string][]workflow, params map[string]int) int {
	next := "in"
	for next != "R" {
		for _, flow := range wf[next] {
			if dest, ok := flow.evaluate(params); ok {
				next = dest
				break
			}
		}
		if next == "A" {
			res := 0
			for _, v := range params {
				res += v
			}
			return res
		}
	}
	return 0
}

func parseWorkflows(s string) []workflow {
	flows := []workflow{}
	for _, flow := range strings.Split(s[1:len(s)-1], ",") {
		parts := strings.Split(flow, ":")
		if len(parts) == 1 {
			flows = append(flows, workflow{func(map[string]int) (string, bool) {
				return parts[0], true
			}})
		} else {
			condition := parts[0]
			target := parts[1]
			flows = append(flows, workflow{func(params map[string]int) (string, bool) {
				if condition[1] == '<' {
					return target, params[string(condition[0])] < toint(condition[2:])
				} else {
					return target, params[string(condition[0])] > toint(condition[2:])
				}
			}})
		}
	}
	return flows
}

func parseParams(s string) map[string]int {
	params := map[string]int{}
	for _, part := range strings.Split(s[1:len(s)-1], ",") {
		kv := strings.Split(part, "=")
		params[kv[0]] = toint(kv[1])
	}
	return params
}
