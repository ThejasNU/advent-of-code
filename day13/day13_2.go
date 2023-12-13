package day13

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day13_2() {
	input := helpers.ReadInputAsString("./day13/day13_input.txt")
	patterns := strings.Split(input, "\n\n")

	ans := FinalAns{
		colsLeft:  0,
		rowsAbove: 0,
	}
	var wg sync.WaitGroup

	for _, pattern := range patterns {
		lines := strings.Split(pattern, "\n")
		wg.Add(1)
		go func(lines []string) {
			defer wg.Done()

			verticalSymmetry := computeVerticalSymmetry2(lines)
			if verticalSymmetry == -1 {
				horizontalSymmetry := computeHorizontalSymmetry2(lines)
				if horizontalSymmetry != -1 {
					ans.lock.Lock()
					ans.rowsAbove += horizontalSymmetry
					ans.lock.Unlock()
				} else {
					log.Fatal("go die")
				}
			} else {
				ans.lock.Lock()
				ans.colsLeft += verticalSymmetry
				ans.lock.Unlock()
			}
		}(lines)
	}

	wg.Wait()
	fmt.Println(ans.colsLeft + (100 * ans.rowsAbove))
}

func computeVerticalSymmetry2(lines []string) int {
	m, n := len(lines), len(lines[0])

	for i := 1; i < n; i++ {
		left, right := i-1, i
		isSymmetric := true
		fixedOneDiff := false
		for left >= 0 && right < n && isSymmetric {
			diffCount := 0
			for j := 0; j < m; j++ {
				if lines[j][left] != lines[j][right] {
					diffCount++
				}
			}
			if diffCount > 1 {
				isSymmetric = false
				break
			} else if diffCount == 1 {
				if fixedOneDiff == false {
					fixedOneDiff = true
				} else {
					isSymmetric = false
					break
				}
			}

			left--
			right++
		}

		if isSymmetric && fixedOneDiff {
			return i
		}
	}

	return -1
}

func computeHorizontalSymmetry2(lines []string) int {
	m, n := len(lines), len(lines[0])

	for i := 1; i < m; i++ {
		up, down := i-1, i
		isSymmetric := true
		fixedOneDiff := false
		for up >= 0 && down < m && isSymmetric {
			diffCount := 0
			for j := 0; j < n; j++ {
				if lines[up][j] != lines[down][j] {
					diffCount++
				}
			}
			if diffCount > 1 {
				isSymmetric = false
				break
			} else if diffCount == 1 {
				if fixedOneDiff == false {
					fixedOneDiff = true
				} else {
					isSymmetric = false
					break
				}
			}
			up--
			down++
		}

		if isSymmetric && fixedOneDiff {
			return i
		}
	}

	return -1
}
