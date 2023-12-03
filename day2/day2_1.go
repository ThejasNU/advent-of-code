package day2

import (
	"fmt"
	"unicode"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day2_1() {

	input := helpers.ReadInput("./day2/day2_input.txt")
	red, green, blue := 12, 13, 14
	ans := 0

	for i := 1; i < 101; i++ {
		s := input[i-1]
		var j int
		if i < 10 {
			j = 8
		} else {
			j = 9
		}

		possible := true

		for j < len(s) {
			num := 0

			for unicode.IsDigit(rune(s[j])) {
				num = num*10 + (int(s[j] - '0'))
				j++
			}
			j++

			if s[j] == 'r' {
				if num > red {
					possible = false
					break
				}
				j += 5
			} else if s[j] == 'g' {
				if num > green {
					possible = false
					break
				}
				j += 7
			} else if s[j] == 'b' {
				if num > blue {
					possible = false
					break
				}
				j += 6
			}
		}

		if possible {
			ans += i
		}
	}

	fmt.Println(ans)
}
