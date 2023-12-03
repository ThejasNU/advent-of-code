package day3

import (
	"fmt"
	"log"
	"strconv"
	"unicode"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day3_1() {
	input := helpers.ReadInput("./day3/day3_input.txt")
	m, n := len(input), len(input[0])
	var numberIdxs [][]int

	symbolIdxs:= getSymbolsIdx(input)

	for _,idx := range symbolIdxs {
		row, col := idx[0], idx[1]

		if col-1 >= 0 {
			if unicode.IsDigit(rune(input[row][col-1])) {
				numberIdxs = append(numberIdxs, []int{row, col - 1})
			}

			if row-1 >= 0 && unicode.IsDigit(rune(input[row-1][col-1])) {
				numberIdxs = append(numberIdxs, []int{row - 1, col - 1})
			}

			if row+1 < m && unicode.IsDigit(rune(input[row+1][col-1])) {
				numberIdxs = append(numberIdxs, []int{row + 1, col - 1})
			}
		}

		if col+1 < n {
			if unicode.IsDigit(rune(input[row][col+1])) {
				numberIdxs = append(numberIdxs, []int{row, col + 1})
			}

			if row-1 >= 0 && unicode.IsDigit(rune(input[row-1][col+1])) {
				numberIdxs = append(numberIdxs, []int{row - 1, col + 1})
			}

			if row+1 < m && unicode.IsDigit(rune(input[row+1][col+1])) {
				numberIdxs = append(numberIdxs, []int{row + 1, col + 1})
			}
		}

		if row+1 < m && unicode.IsDigit(rune(input[row+1][col])) {
			numberIdxs = append(numberIdxs, []int{row + 1, col})
		}

		if row-1 >= 0 && unicode.IsDigit(rune(input[row-1][col])) {
			numberIdxs = append(numberIdxs, []int{row - 1, col})
		}

	}
	getNumbers(input, numberIdxs)
}

func getSymbolsIdx(matrix []string) [][]int {
	var symbolIdxs [][]int
	for i, row := range matrix {
		for j, char := range row {
			asciiVal := int(char)
			if asciiVal != 46 && (asciiVal > 57 || asciiVal < 48) {
				symbolIdxs = append(symbolIdxs, []int{i, j})
			}
		}
	}
	return symbolIdxs
}

func getNumbers(matrix []string, indexes [][]int) {
	sum := 0
	rows,cols:=len(matrix),len(matrix[0])
	vis := make([][]int, rows)
	for i := 0; i < rows; i++ {
		vis[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			vis[i][j] = 0
		}
	}
	for _, idx := range indexes {
		row, col := idx[0], idx[1]
		if vis[row][col]==1{
			continue
		}
		vis[row][col]=1
		l, r := col-1, col+1
		for l >= 0 && unicode.IsDigit(rune(matrix[row][l])) {
			vis[row][l]=1
			l--
		}
		
		for r < len(matrix[0]) && unicode.IsDigit(rune(matrix[row][r])) {
			vis[row][r]=1
			r++
		}

		rowStr := matrix[row]
		numStr := rowStr[l+1 : r]
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}
		sum += num
	}

	fmt.Println(sum)
}
