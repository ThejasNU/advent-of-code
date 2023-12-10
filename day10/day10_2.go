package day10

import (
	"container/heap"
	"fmt"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day10_2() {
	input := helpers.ReadInput("./day10/day10_input.txt")
	m, n := len(input), len(input[0])
	startIdx := make([]int, 2)
	found := false
	for i, row := range input {
		for j, col := range row {
			if rune(col) == 'S' {
				found = true
				startIdx[0] = i
				startIdx[1] = j
				break
			}
		}
		if found == true {
			break
		}
	}
	pipeVisited := make([][]bool, m)
	for i := range pipeVisited {
		pipeVisited[i] = make([]bool, n)
	}
	pipeVisited[startIdx[0]][startIdx[1]] = true
	pq := make(helpers.PriorityQueue, 0)
	heap.Init(&pq)

	if startIdx[0]-1 >= 0 {
		char := rune(input[startIdx[0]-1][startIdx[1]])
		if char == '|' || char == '7' || char == 'F' {
			heap.Push(&pq, &helpers.Item{
				Value:    []int{startIdx[0] - 1, startIdx[1]},
				Priority: 1,
			})
			pipeVisited[startIdx[0]-1][startIdx[1]] = true
		}
	}
	if startIdx[0]+1 < m {
		char := rune(input[startIdx[0]+1][startIdx[1]])
		if char == '|' || char == 'L' || char == 'J' {
			heap.Push(&pq, &helpers.Item{
				Value:    []int{startIdx[0] + 1, startIdx[1]},
				Priority: 1,
			})
			pipeVisited[startIdx[0]+1][startIdx[1]] = true
		}
	}
	if startIdx[1]-1 >= 0 {
		char := rune(input[startIdx[0]][startIdx[1]-1])
		if char == '-' || char == 'L' || char == 'F' {
			heap.Push(&pq, &helpers.Item{
				Value:    []int{startIdx[0], startIdx[1] - 1},
				Priority: 1,
			})
			pipeVisited[startIdx[0]][startIdx[1]-1] = true
		}
	}
	if startIdx[1]+1 < n {
		char := rune(input[startIdx[0]][startIdx[1]+1])
		if char == 'J' || char == '-' || char == '7' {
			heap.Push(&pq, &helpers.Item{
				Value:    []int{startIdx[0], startIdx[1] + 1},
				Priority: 1,
			})
			pipeVisited[startIdx[0]][startIdx[1]+1] = true
		}
	}

	for pq.Len() > 0 {
		ele := heap.Pop(&pq).(*helpers.Item)
		i, j := ele.Value[0], ele.Value[1]
		steps := ele.Priority

		curChar := rune(input[i][j])
		switch curChar {
		case '|':
			if i+1 < m && !pipeVisited[i+1][j] {
				nextChar := rune(input[i+1][j])
				if nextChar == 'L' || nextChar == 'J' || nextChar == '|' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i + 1, j},
						Priority: steps + 1,
					})
					pipeVisited[i+1][j] = true
				}
			}
			if i-1 >= 0 && !pipeVisited[i-1][j] {
				nextChar := rune(input[i-1][j])
				if nextChar == 'F' || nextChar == '7' || nextChar == '|' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i - 1, j},
						Priority: steps + 1,
					})
					pipeVisited[i-1][j] = true
				}
			}
		case '-':
			if j-1 >= 0 && !pipeVisited[i][j-1] {
				nextChar := rune(input[i][j-1])
				if nextChar == 'L' || nextChar == 'F' || nextChar == '-' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i, j - 1},
						Priority: steps + 1,
					})
					pipeVisited[i][j-1] = true
				}
			}
			if j+1 < n && !pipeVisited[i][j+1] {
				nextChar := rune(input[i][j+1])
				if nextChar == '7' || nextChar == 'J' || nextChar == '-' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i, j + 1},
						Priority: steps + 1,
					})
					pipeVisited[i][j+1] = true
				}
			}
		case 'L':
			if j+1 < n && !pipeVisited[i][j+1] {
				nextChar := rune(input[i][j+1])
				if nextChar == '7' || nextChar == 'J' || nextChar == '-' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i, j + 1},
						Priority: steps + 1,
					})
					pipeVisited[i][j+1] = true
				}
			}
			if i-1 >= 0 && !pipeVisited[i-1][j] {
				nextChar := rune(input[i-1][j])
				if nextChar == '7' || nextChar == 'F' || nextChar == '|' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i - 1, j},
						Priority: steps + 1,
					})
					pipeVisited[i-1][j] = true
				}
			}
		case 'J':
			if j-1 >= 0 && !pipeVisited[i][j-1] {
				nextChar := rune(input[i][j-1])
				if nextChar == 'L' || nextChar == 'F' || nextChar == '-' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i, j - 1},
						Priority: steps + 1,
					})
					pipeVisited[i][j-1] = true
				}
			}
			if i-1 >= 0 && !pipeVisited[i-1][j] {
				nextChar := rune(input[i-1][j])
				if nextChar == '7' || nextChar == 'F' || nextChar == '|' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i - 1, j},
						Priority: steps + 1,
					})
					pipeVisited[i-1][j] = true
				}
			}
		case '7':
			if j-1 >= 0 && !pipeVisited[i][j-1] {
				nextChar := rune(input[i][j-1])
				if nextChar == 'L' || nextChar == 'F' || nextChar == '-' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i, j - 1},
						Priority: steps + 1,
					})
					pipeVisited[i][j-1] = true
				}
			}
			if i+1 < m && !pipeVisited[i+1][j] {
				nextChar := rune(input[i+1][j])
				if nextChar == 'L' || nextChar == 'J' || nextChar == '|' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i + 1, j},
						Priority: steps + 1,
					})
					pipeVisited[i+1][j] = true
				}
			}
		case 'F':
			if i+1 < m && !pipeVisited[i+1][j] {
				nextChar := rune(input[i+1][j])
				if nextChar == 'L' || nextChar == 'J' || nextChar == '|' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i + 1, j},
						Priority: steps + 1,
					})
					pipeVisited[i+1][j] = true
				}
			}
			if j+1 < n && !pipeVisited[i][j+1] {
				nextChar := rune(input[i][j+1])
				if nextChar == '7' || nextChar == 'J' || nextChar == '-' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i, j + 1},
						Priority: steps + 1,
					})
					pipeVisited[i][j+1] = true
				}
			}
		}
	}

	/* Imagine a circle if a point is inside, if we move diagonally it will cross the border only once
	If it is outside, it will cross border 0 times or 2 times, this trick is applicable for all closed shapes
	Edge case is tangent,7 and L means we are tangent to the closed path, so better to ignore it instead of writing new condition and writing 2
	*/
	insideCount := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !pipeVisited[i][j] {
				borderCrosses := 0
				r, c := i, j

				for r < m && c < n {
					curTile := rune(input[r][c])
					if pipeVisited[r][c] && curTile != '7' && curTile != 'L' {
						borderCrosses++
					}
					r++
					c++
				}

				if borderCrosses%2 == 1 {
					insideCount++
				}
			}
		}
	}

	fmt.Println(insideCount)
}
