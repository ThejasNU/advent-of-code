package day12

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day12_2() {
	input := helpers.ReadInput("./day12/day12_input.txt")
	var wg sync.WaitGroup
	ans := FinalAns{
		val: 0,
	}
	for _, s := range input {
		springs, groups := getUnfoldedSpringsAndGroups(s)

		wg.Add(1)
		go func(springs string, groups []int) {
			defer wg.Done()
			//springsIdx,groupsIdx,currentCount
			dp := make(map[int]map[int]map[int]int)
			res := computeNumArrangements(springs, groups, 0, 0, 0, dp)
			ans.lock.Lock()
			ans.val += res
			ans.lock.Unlock()
		}(springs, groups)
	}
	wg.Wait()
	fmt.Println(ans.val)
}

func getUnfoldedSpringsAndGroups(s string) (string, []int) {
	res := strings.Split(s, " ")
	springs := res[0]
	unfoldedSprings := springs + "?" + springs + "?" + springs + "?" + springs + "?" + springs
	groupStrs := strings.Split(res[1], ",")
	groups := make([]int, len(groupStrs)*5)
	n := len(groupStrs)
	for idx, groupStr := range groupStrs {
		groupVal, _ := strconv.ParseInt(groupStr, 10, 64)
		groups[idx] = int(groupVal)
		groups[idx+n] = int(groupVal)
		groups[idx+(2*n)] = int(groupVal)
		groups[idx+(3*n)] = int(groupVal)
		groups[idx+(4*n)] = int(groupVal)
	}

	return unfoldedSprings, groups
}
