package day11

import (
	"fmt"
	"math"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day11_1() {
	input := helpers.ReadInput("./day11/day11_input.txt")
	galaxyIdxs := getGalaxyIdxs(input)
	emptyRows := getEmptyRows(input)
	emptyCols := getEmptyCols(input)

	ans := 0
	n := len(galaxyIdxs)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := int(math.Abs(float64(galaxyIdxs[j][0]-galaxyIdxs[i][0]))) + int(math.Abs(float64(galaxyIdxs[j][1]-galaxyIdxs[i][1])))
			minR := int(math.Min(float64(galaxyIdxs[i][0]), float64(galaxyIdxs[j][0])))
			maxR := int(math.Max(float64(galaxyIdxs[i][0]), float64(galaxyIdxs[j][0])))
			minC := int(math.Min(float64(galaxyIdxs[i][1]), float64(galaxyIdxs[j][1])))
			maxC := int(math.Max(float64(galaxyIdxs[i][1]), float64(galaxyIdxs[j][1])))

			for _, emptyRow := range emptyRows {
				if minR <= emptyRow && emptyRow <= maxR {
					dist++
				}
			}
			for _, emptyCol := range emptyCols {
				if minC <= emptyCol && emptyCol <= maxC {
					dist++
				}
			}
			ans += dist
		}
	}
	fmt.Println(ans)
}

func getGalaxyIdxs(input []string) [][]int {
	var galaxyIdxs [][]int

	for i, row := range input {
		for j, char := range row {
			if char == '#' {
				galaxyIdxs = append(galaxyIdxs, []int{i, j})
			}
		}
	}

	return galaxyIdxs
}

func getEmptyRows(input []string) []int {
	m, n := len(input), len(input[0])
	var emptyRows []int
	for i := 0; i < m; i++ {
		isEmpty := true
		for j := 0; j < n; j++ {
			if rune(input[i][j]) != '.' {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			emptyRows = append(emptyRows, i)
		}
	}
	return emptyRows
}

func getEmptyCols(input []string) []int {
	m, n := len(input), len(input[0])
	var emptyCols []int
	for i := 0; i < n; i++ {
		isEmpty := true
		for j := 0; j < m; j++ {
			if rune(input[j][i]) != '.' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyCols = append(emptyCols, i)
		}
	}
	return emptyCols
}
