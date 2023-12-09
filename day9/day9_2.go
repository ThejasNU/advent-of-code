package day9

import (
	"fmt"
	"regexp"
	"strconv"
	"sync"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day9_2() {
	input := helpers.ReadInput("./day9/day9_input.txt")
	var ans FinalAns = FinalAns{
		val: 0,
	}
	regex, _ := regexp.Compile("-?[0-9]+")
	var wg sync.WaitGroup

	for _, s := range input {
		numStrs := regex.FindAll([]byte(s), -1)
		n := len(numStrs)
		sequence := make([]int64, n)

		for i := n - 1; i >= 0; i-- {
			num, _ := strconv.ParseInt(string(numStrs[i]), 10, 64)
			sequence[n-1-i] = num
		}

		wg.Add(1)
		go func(sequence []int64) {
			defer wg.Done()
			ans.computeNextValue(sequence)
		}(sequence)
	}

	wg.Wait()
	fmt.Println(ans.val)
}
