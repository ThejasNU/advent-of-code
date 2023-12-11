package day11

import (
	"fmt"
	"math"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day11_2() {
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
					dist += 999999
				}
			}
			for _, emptyCol := range emptyCols {
				if minC <= emptyCol && emptyCol <= maxC {
					dist += 999999
				}
			}
			ans += dist
		}
	}
	fmt.Println(ans)
}
