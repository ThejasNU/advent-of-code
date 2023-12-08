package day8

import (
	"fmt"
	"regexp"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day8_1() {
	input := helpers.ReadInput("./day8/day8_input.txt")

	instructions := input[0]

	pathsMap := getPathMap(input[2:])
	curNode := "AAA"
	i := 0
	steps := 0
	for curNode != "ZZZ" {
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

	fmt.Println(steps)
}

func getPathMap(inputs []string) map[string][]string {
	regex, _ := regexp.Compile("[A-Z]+")
	m := make(map[string][]string)
	for _, s := range inputs {
		nodes := regex.FindAll([]byte(s), -1)
		m[string(nodes[0])] = []string{string(nodes[1]), string(nodes[2])}
	}
	return m
}
