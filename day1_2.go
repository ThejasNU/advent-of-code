package main

import (
	"fmt"
	"unicode"
)

func Day1_2(){
	input := ReadInput("day1_input.txt")
	ans :=0
	for _,s:=range input{
		n := len(s)
		var nums []int
		i := 0
		for i<n{
			if unicode.IsDigit(rune(s[i])){
				nums=append(nums,int(s[i]-'0'))
			}

			if i+2<n{
				num:=s[i:i+3]

				if num=="one"{
					nums=append(nums, 1)
				} else if num=="two"{
					nums=append(nums, 2)
				} else if num=="six"{
					nums=append(nums, 6)
				}
			}

			if i+3<n{
				num := s[i:i+4]

				if num=="four"{
					nums=append(nums, 4)
				} else if num=="five"{
					nums=append(nums, 5)
				} else if num=="nine"{
					nums=append(nums, 9)
				}
			}

			if i+4<n{
				num:=s[i:i+5]

				if num=="three"{
					nums=append(nums, 3)
				} else if num=="seven"{
					nums=append(nums, 7)
				} else if num=="eight"{
					nums=append(nums, 8)
				}
			}

			i++
		}
		ans+= nums[0]*10 + nums[len(nums)-1]
	}

	fmt.Println(ans)
}