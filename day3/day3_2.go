package day3

import (
	"fmt"
	"log"
	"strconv"
	"unicode"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day3_2() {
	input := helpers.ReadInput("./day3/day3_input.txt")
	starIdxs := getStarIdxs(input)
	gearsIdxs := getGears(input, starIdxs)
	printAns(input,gearsIdxs)
}

func getStarIdxs(matrix []string) [][]int {
	var starIdxs [][]int
	for i, row := range matrix {
		for j, char := range row {
			if char == '*' {
				starIdxs = append(starIdxs, []int{i, j})
			}
		}
	}
	return starIdxs
}

func getGears(input []string, starIdxs [][]int) [][][]int {
	m, n := len(input), len(input[0])
	var gearIdxs [][][]int

	for _, idx := range starIdxs {
		row, col := idx[0], idx[1]
		var gear [][]int
		if col-1 >= 0 {
			if unicode.IsDigit(rune(input[row][col-1])) {
				gear = append(gear, []int{row, col - 1})
			}

			if row-1 >= 0 && unicode.IsDigit(rune(input[row-1][col-1])) {
				gear = append(gear, []int{row - 1, col - 1})
			}

			if row+1 < m && unicode.IsDigit(rune(input[row+1][col-1])) {
				gear = append(gear, []int{row + 1, col - 1})
			}
		}

		if col+1 < n {
			if unicode.IsDigit(rune(input[row][col+1])) {
				gear = append(gear, []int{row, col + 1})
			}

			if row-1 >= 0 && unicode.IsDigit(rune(input[row-1][col+1])) {
				gear = append(gear, []int{row - 1, col + 1})
			}

			if row+1 < m && unicode.IsDigit(rune(input[row+1][col+1])) {
				gear = append(gear, []int{row + 1, col + 1})
			}
		}

		if row+1 < m && unicode.IsDigit(rune(input[row+1][col])) {
			gear = append(gear, []int{row + 1, col})
		}

		if row-1 >= 0 && unicode.IsDigit(rune(input[row-1][col])) {
			gear = append(gear, []int{row - 1, col})
		}

		gearIdxs = append(gearIdxs, gear)
	}

	return gearIdxs
}

func printAns(input []string, gearIdxs [][][]int) {
	ans := 0
	m, n := len(input), len(input[0])
	vis := make([][]int, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			vis[i][j] = 0
		}
	}

	for _, gear := range gearIdxs {
		var nums []int
		for _, idx := range gear {
			row, col := idx[0], idx[1]
			if vis[row][col] == 1 {
				continue
			}
			vis[row][col] = 1
			l, r := col-1, col+1
			for l >= 0 && unicode.IsDigit(rune(input[row][l])) {
				vis[row][l] = 1
				l--
			}

			for r < n && unicode.IsDigit(rune(input[row][r])) {
				vis[row][r] = 1
				r++
			}
			rowStr := input[row]
			numStr := rowStr[l+1 : r]
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, num)
		}
		if len(nums) == 2 {
			ans += nums[0] * nums[1]
		}
	}

	fmt.Println(ans)
}
