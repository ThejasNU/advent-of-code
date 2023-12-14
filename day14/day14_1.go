package day14

import (
	"fmt"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day14_1() {
	input := helpers.ReadInput("./day14/day14_input.txt")
	m, n := len(input), len(input[0])
	ans := 0
	for col := 0; col < n; col++ {
		count := 0
		for row := m - 1; row >= 0; row-- {
			if rune(input[row][col]) == 'O' {
				count++
			} else if rune(input[row][col]) == '#' && count > 0 {
				startDist := m - row - 1
				for count > 0 {
					ans += startDist
					startDist--
					count--
				}
			}
		}

		if count > 0 {
			startDist := m
			for count > 0 {
				ans += startDist
				startDist--
				count--
			}
		}
	}
	fmt.Println(ans)
}
