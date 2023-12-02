package main

import (
	"fmt"
	"unicode"
)

func Day2_2() {

	input := ReadInput("day2_input.txt")
	ans := 0

	for i := 1; i < 101; i++ {
		red,green,blue := 0,0,0
		s := input[i-1]
		var j int
		if i < 10 {
			j = 8
		} else {
			j = 9
		}

		for j < len(s) {
			num := 0

			for unicode.IsDigit(rune(s[j])) {
				num = num*10 + (int(s[j] - '0'))
				j++
			}
			j++

			switch s[j]{
			case 'r':
				red=max(red,num)
				j+=5
			case 'g':
				green=max(green,num)
				j+=7
			case 'b':
				blue=max(blue,num)
				j+=6
			}
		}

		ans+=red*green*blue
	}

	fmt.Println(ans)
}
