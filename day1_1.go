package main

import (
	"fmt"
	"unicode"
)

func Day1_1(){
	input := ReadInput("day1_input.txt")
	ans:=0

	for _,s:=range input{
		var nums []int
		for _,v:=range s{
			if unicode.IsDigit(v){
				nums = append(nums, int(v)-48)
			}
		}

		ans+= nums[0]*10+nums[len(nums)-1]
	}

	fmt.Println(ans)
}