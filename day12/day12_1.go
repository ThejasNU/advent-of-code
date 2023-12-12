package day12

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day12_1() {
	input := helpers.ReadInput("./day12/day12_input.txt")
	var wg sync.WaitGroup
	ans := FinalAns{
		val: 0,
	}
	for _, s := range input {
		springs, groups := getSpringsAndGroups(s)
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

type FinalAns struct {
	val  int
	lock sync.Mutex
}

func getSpringsAndGroups(s string) (string, []int) {
	res := strings.Split(s, " ")
	springs := res[0]
	groupStrs := strings.Split(res[1], ",")
	groups := make([]int, len(groupStrs))

	for idx, groupStr := range groupStrs {
		groupVal, _ := strconv.ParseInt(groupStr, 10, 64)
		groups[idx] = int(groupVal)
	}

	return springs, groups
}

func computeNumArrangements(springs string, groups []int, springsIdx int, groupsIdx int, currentCount int, dp map[int]map[int]map[int]int) int {
	if springsIndexMap, ok := dp[springsIdx]; ok {
		if groupsIndexMap, okay := springsIndexMap[groupsIdx]; okay {
			if val, okie := groupsIndexMap[currentCount]; okie {
				return val
			}
		}
	}

	if springsIdx == len(springs) {
		if groupsIdx == len(groups) && currentCount == 0 {
			return 1
		} else if groupsIdx == len(groups)-1 && groups[groupsIdx] == currentCount {
			return 1
		} else {
			return 0
		}
	}

	res := 0

	if rune(springs[springsIdx]) == '.' {
		if currentCount == 0 {
			res += computeNumArrangements(springs, groups, springsIdx+1, groupsIdx, 0, dp)
		} else if currentCount > 0 && groupsIdx < len(groups) && groups[groupsIdx] == currentCount {
			res += computeNumArrangements(springs, groups, springsIdx+1, groupsIdx+1, 0, dp)
		}
	} else if rune(springs[springsIdx]) == '#' {
		res += computeNumArrangements(springs, groups, springsIdx+1, groupsIdx, currentCount+1, dp)
	} else {
		//take ? as .
		if currentCount == 0 {
			res += computeNumArrangements(springs, groups, springsIdx+1, groupsIdx, 0, dp)
		} else if currentCount > 0 && groupsIdx < len(groups) && groups[groupsIdx] == currentCount {
			res += computeNumArrangements(springs, groups, springsIdx+1, groupsIdx+1, 0, dp)
		}

		//take ? as #
		res += computeNumArrangements(springs, groups, springsIdx+1, groupsIdx, currentCount+1, dp)
	}

	if springsIndexMap, ok := dp[springsIdx]; ok {
		if _, okay := springsIndexMap[groupsIdx]; okay {
			dp[springsIdx][groupsIdx][currentCount] = res
		} else {
			dp[springsIdx][groupsIdx] = make(map[int]int)
			dp[springsIdx][groupsIdx][currentCount] = res
		}
	} else {
		dp[springsIdx] = make(map[int]map[int]int)
		dp[springsIdx][groupsIdx] = make(map[int]int)
		dp[springsIdx][groupsIdx][currentCount] = res
	}

	return res
}
