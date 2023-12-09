package day9

import (
	"fmt"
	"regexp"
	"strconv"
	"sync"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day9_1() {
	input := helpers.ReadInput("./day9/day9_input.txt")
	var ans FinalAns = FinalAns{
		val: 0,
	}
	regex, _ := regexp.Compile("-?[0-9]+")
	var wg sync.WaitGroup

	for _, s := range input {
		numStrs := regex.FindAll([]byte(s), -1)
		sequence := make([]int64, len(numStrs))

		for idx, numBytes := range numStrs {
			num, _ := strconv.ParseInt(string(numBytes), 10, 64)
			sequence[idx] = num
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

type FinalAns struct {
	val  int64
	lock sync.Mutex
}

func (ans *FinalAns) computeNextValue(sequence []int64) {
	var nextVal int64 = 0
	for {
		nextVal += sequence[len(sequence)-1]
		differencesSequence := make([]int64, len(sequence)-1)
		numZeros := 0
		for i := 1; i < len(sequence); i++ {
			differencesSequence[i-1] = sequence[i] - sequence[i-1]
			if differencesSequence[i-1] == 0 {
				numZeros++
			}
		}
		if numZeros == len(differencesSequence) {
			break
		}
		sequence = differencesSequence
	}

	ans.lock.Lock()
	ans.val += nextVal
	ans.lock.Unlock()
}
