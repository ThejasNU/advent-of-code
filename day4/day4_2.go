package day4

import (
	"fmt"
	"strings"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day4_2() {
	input := helpers.ReadInput("./day4/day4_input.txt")
	var pointsList []int
	for _, s := range input {
		winMap := getWinMap(s[10:39])
		count := getCount(s[42:], winMap)
		pointsList = append(pointsList, count)
	}
	printAns(pointsList, len(input))
}

func getCount(nums string, winMap map[string]int) int {
	count := 0
	numsList := strings.Split(nums, " ")

	for _, num := range numsList {
		if _, ok := winMap[num]; ok {
			count++
		}
	}

	return count
}

func printAns(pointsList []int, numCards int) {
	cardsMap := make(map[int]int)
	for i := 1; i <= numCards; i++ {
		cardsMap[i] = 1
	}

	for idx, points := range pointsList {
		cardNum := idx + 1
		numCardsInMap, _ := cardsMap[cardNum]
		for i := 1; i <= points; i++ {
			if val, ok := cardsMap[cardNum+i]; ok {
				cardsMap[cardNum+i] = val + numCardsInMap
			} else {
				cardsMap[cardNum+i] = 1
			}
		}
	}

	ans := 0
	for _, val := range cardsMap {
		ans += val
	}

	fmt.Println(ans)
}
