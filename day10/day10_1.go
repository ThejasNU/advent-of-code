package day10

import (
	"container/heap"
	"fmt"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day10_1() {
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
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	visited[startIdx[0]][startIdx[1]] = true
	pq := make(helpers.PriorityQueue, 0)
	heap.Init(&pq)

	if startIdx[0]-1 >= 0 {
		char := rune(input[startIdx[0]-1][startIdx[1]])
		if char == '|' || char == '7' || char == 'F' {
			heap.Push(&pq, &helpers.Item{
				Value:    []int{startIdx[0] - 1, startIdx[1]},
				Priority: 1,
			})
			visited[startIdx[0]-1][startIdx[1]] = true
		}
	}
	if startIdx[0]+1 < m {
		char := rune(input[startIdx[0]+1][startIdx[1]])
		if char == '|' || char == 'L' || char == 'J' {
			heap.Push(&pq, &helpers.Item{
				Value:    []int{startIdx[0] + 1, startIdx[1]},
				Priority: 1,
			})
			visited[startIdx[0]+1][startIdx[1]] = true
		}
	}
	if startIdx[1]-1 >= 0 {
		char := rune(input[startIdx[0]][startIdx[1]-1])
		if char == '-' || char == 'L' || char == 'F' {
			heap.Push(&pq, &helpers.Item{
				Value:    []int{startIdx[0], startIdx[1] - 1},
				Priority: 1,
			})
			visited[startIdx[0]][startIdx[1]-1] = true
		}
	}
	if startIdx[1]+1 < n {
		char := rune(input[startIdx[0]][startIdx[1]+1])
		if char == 'J' || char == '-' || char == '7' {
			heap.Push(&pq, &helpers.Item{
				Value:    []int{startIdx[0], startIdx[1] + 1},
				Priority: 1,
			})
			visited[startIdx[0]][startIdx[1]+1] = true
		}
	}

	ans := 0
	for pq.Len() > 0 {
		ele := heap.Pop(&pq).(*helpers.Item)
		i, j := ele.Value[0], ele.Value[1]
		steps := ele.Priority
		if steps > ans {
			ans = steps
		}

		curChar := rune(input[i][j])
		switch curChar {
		case '|':
			if i+1 < m && !visited[i+1][j] {
				nextChar := rune(input[i+1][j])
				if nextChar == 'L' || nextChar == 'J' || nextChar == '|' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i + 1, j},
						Priority: steps + 1,
					})
					visited[i+1][j] = true
				}
			}
			if i-1 >= 0 && !visited[i-1][j] {
				nextChar := rune(input[i-1][j])
				if nextChar == 'F' || nextChar == '7' || nextChar == '|' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i - 1, j},
						Priority: steps + 1,
					})
					visited[i-1][j] = true
				}
			}
		case '-':
			if j-1 >= 0 && !visited[i][j-1] {
				nextChar := rune(input[i][j-1])
				if nextChar == 'L' || nextChar == 'F' || nextChar == '-' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i, j - 1},
						Priority: steps + 1,
					})
					visited[i][j-1] = true
				}
			}
			if j+1 < n && !visited[i][j+1] {
				nextChar := rune(input[i][j+1])
				if nextChar == '7' || nextChar == 'J' || nextChar == '-' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i, j + 1},
						Priority: steps + 1,
					})
					visited[i][j+1] = true
				}
			}
		case 'L':
			if j+1 < n && !visited[i][j+1] {
				nextChar := rune(input[i][j+1])
				if nextChar == '7' || nextChar == 'J' || nextChar == '-' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i, j + 1},
						Priority: steps + 1,
					})
					visited[i][j+1] = true
				}
			}
			if i-1 >= 0 && !visited[i-1][j] {
				nextChar := rune(input[i-1][j])
				if nextChar == '7' || nextChar == 'F' || nextChar == '|' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i - 1, j},
						Priority: steps + 1,
					})
					visited[i-1][j] = true
				}
			}
		case 'J':
			if j-1 >= 0 && !visited[i][j-1] {
				nextChar := rune(input[i][j-1])
				if nextChar == 'L' || nextChar == 'F' || nextChar == '-' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i, j - 1},
						Priority: steps + 1,
					})
					visited[i][j-1] = true
				}
			}
			if i-1 >= 0 && !visited[i-1][j] {
				nextChar := rune(input[i-1][j])
				if nextChar == '7' || nextChar == 'F' || nextChar == '|' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i - 1, j},
						Priority: steps + 1,
					})
					visited[i-1][j] = true
				}
			}
		case '7':
			if j-1 >= 0 && !visited[i][j-1] {
				nextChar := rune(input[i][j-1])
				if nextChar == 'L' || nextChar == 'F' || nextChar == '-' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i, j - 1},
						Priority: steps + 1,
					})
					visited[i][j-1] = true
				}
			}
			if i+1 < m && !visited[i+1][j] {
				nextChar := rune(input[i+1][j])
				if nextChar == 'L' || nextChar == 'J' || nextChar == '|' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i + 1, j},
						Priority: steps + 1,
					})
					visited[i+1][j] = true
				}
			}
		case 'F':
			if i+1 < m && !visited[i+1][j] {
				nextChar := rune(input[i+1][j])
				if nextChar == 'L' || nextChar == 'J' || nextChar == '|' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i + 1, j},
						Priority: steps + 1,
					})
					visited[i+1][j] = true
				}
			}
			if j+1 < n && !visited[i][j+1] {
				nextChar := rune(input[i][j+1])
				if nextChar == '7' || nextChar == 'J' || nextChar == '-' {
					heap.Push(&pq, &helpers.Item{
						Value:    []int{i, j + 1},
						Priority: steps + 1,
					})
					visited[i][j+1] = true
				}
			}
		}
	}
	fmt.Println(ans)
}
