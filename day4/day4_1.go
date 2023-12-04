package day4

import (
	"fmt"
	"math"

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
	for i:=0;i<len(nums);i+=3{
		if nums[i]==' '{
			m[string(nums[i+1])]=1
		} else{
			m[string(nums[i])+string(nums[i+1])]=1
		}
	}
	return m
}

func getPoints(nums string,winMap map[string]int) int{
	count:=0
	for i:=0;i<len(nums);i+=3{
		var num string
		if nums[i]==' '{
			num=string(nums[i+1])
		} else{
			num=string(nums[i])+string(nums[i+1])
		}

		if _,ok:=winMap[num];ok{
			count++
		}
	}

	return int(math.Pow(2,float64(count-1)))
}