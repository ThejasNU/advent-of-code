package day8

import (
	"fmt"
	"sync"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day8_2() {
	input := helpers.ReadInput("./day8/day8_input.txt")
	instructions := input[0]
	pathsMap := getPathMap(input[2:])
	var ans FinalAns
	ans.val = 1

	var wg sync.WaitGroup

	for key := range pathsMap {
		if rune(key[2]) == 'A' {
			wg.Add(1)
			go func(startNode string) {
				defer wg.Done()
				ans.minStepsRequired(startNode, pathsMap, instructions)
			}(key)
		}
	}
	wg.Wait()
	fmt.Println(ans.val)
}

type FinalAns struct {
	val  int
	lock sync.Mutex
}

func (ans *FinalAns) minStepsRequired(startNode string, pathsMap map[string][]string, instructions string) {
	curNode := startNode
	i := 0
	steps := 0
	for rune(curNode[2]) != 'Z' {
		var nextNode string

		if rune(instructions[i]) == 'L' {
			nextNode = pathsMap[curNode][0]
		} else {
			nextNode = pathsMap[curNode][1]
		}

		curNode = nextNode
		i = (i + 1) % len(instructions)
		steps++
	}

	ans.lock.Lock()
	temp := lcm(ans.val, steps)
	ans.val = temp
	ans.lock.Unlock()
}

func lcm(a, b int) int {
	A, B := a, b
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	//here a is GCD of a,b
	return (A * B) / a
}
