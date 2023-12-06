package day6

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day6_2(){
	input:=helpers.ReadInput("./day6/day6_input.txt")
	
	regex,_:=regexp.Compile("[0-9]+")
	
	timeStrs:=regex.FindAll([]byte(strings.Split(input[0], ":")[1]),-1)
	finalTimeStr:=""
	for _,val:=range timeStrs{
		finalTimeStr+=string(val)
	}
	time,_:=strconv.ParseInt(finalTimeStr,10,64)

	distStrs:=regex.FindAll([]byte(strings.Split(input[1], ":")[1]),-1)
	finalDistStr:=""
	for _,val:=range distStrs{
		finalDistStr+=string(val)
	}
	dist,_:=strconv.ParseInt(finalDistStr,10,64)

	fmt.Println(getNumPossibleInConstantTime(time,dist))
}