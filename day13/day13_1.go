package day13

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day13_1() {
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

			verticalSymmetry := computeVerticalSymmetry(lines)
			if verticalSymmetry == -1 {
				horizontalSymmetry := computeHorizontalSymmetry(lines)
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

type FinalAns struct {
	colsLeft  int
	rowsAbove int
	lock      sync.Mutex
}

func computeVerticalSymmetry(lines []string) int {
	m, n := len(lines), len(lines[0])

	for i := 1; i < n; i++ {
		left, right := i-1, i
		isSymmetric := true
		for left >= 0 && right < n && isSymmetric {
			for j := 0; j < m; j++ {
				if lines[j][left] != lines[j][right] {
					isSymmetric = false
					break
				}
			}
			left--
			right++
		}

		if isSymmetric {
			return i
		}
	}

	return -1
}

func computeHorizontalSymmetry(lines []string) int {
	m, n := len(lines), len(lines[0])

	for i := 1; i < m; i++ {
		up, down := i-1, i
		isSymmetric := true
		for up >= 0 && down < m && isSymmetric {
			for j := 0; j < n; j++ {
				if lines[up][j] != lines[down][j] {
					isSymmetric = false
					break
				}
			}
			up--
			down++
		}

		if isSymmetric {
			return i
		}
	}

	return -1
}
