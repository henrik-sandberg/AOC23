package main

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func min(n ...int) int {
	min := n[0]
	for _, v := range n[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func max(n ...int) int {
	max := n[0]
	for _, v := range n[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func multiply(n ...int) int {
	product := n[0]
	for _, v := range n[1:] {
		product *= v
	}
	return product
}
