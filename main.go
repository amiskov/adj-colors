package main

type color string

func longestAdjColors(colors [][]color) int {
	if len(colors) == 0 {
		return 0
	}
	lookup := makeLookupTable(colors)
	res := 0
	for row := 0; row < len(colors); row++ {
		for col := 0; col < len(colors[row]); col++ {
			if !lookup[row][col] {
				res = max(res, countAdjacent(colors, lookup, colors[row][col], row, col))
			}
		}
	}
	return res
}

func countAdjacent(colors [][]color, lookup [][]bool, expected color, row, col int) int {
	if row < 0 || row >= len(colors) {
		return 0
	}
	if col < 0 || col >= len(colors[row]) {
		return 0
	}
	if lookup[row][col] || colors[row][col] != expected {
		return 0
	}
	lookup[row][col] = true
	count := 1
	count += countAdjacent(colors, lookup, expected, row, col-1)
	count += countAdjacent(colors, lookup, expected, row, col+1)
	count += countAdjacent(colors, lookup, expected, row-1, col)
	count += countAdjacent(colors, lookup, expected, row+1, col)
	return count
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func makeLookupTable(colors [][]color) [][]bool {
	lines, cols := len(colors), len(colors[0])
	bools := make([][]bool, lines)
	for l := range bools {
		bools[l] = make([]bool, cols)
	}
	return bools
}

func main() {}
