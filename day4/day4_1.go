package day4

import (
	"fmt"
	"math"
	"strings"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day4_1(){
	input:=helpers.ReadInput("./day4/day4_input.txt")
	ans:=0
	for _,s:=range input{
		winMap:=getWinMap(s[10:39])
		points:=getPoints(s[42:],winMap)
		ans+=points
	}
	fmt.Println(ans)
}

func getWinMap(nums string) map[string] int{
	m:=make(map[string] int)
	numsList := strings.Split(nums," ")
	for _,num:=range numsList{
		if num!="" && num!=" " {
			m[num]=1
		}
	}
	return m
}

func getPoints(nums string,winMap map[string]int) int{
	count:=0
	numsList := strings.Split(nums," ")
	
	for _,num := range numsList{
		if _,ok:=winMap[num];ok{
			count++
		}
	}

	return int(math.Pow(2,float64(count-1)))
}