package day6

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day6_1(){
	input:=helpers.ReadInput("./day6/day6_input.txt")

	regex,_:=regexp.Compile("[0-9]+")

	timeStrs:=regex.FindAll([]byte(strings.Split(input[0], ":")[1]),-1)
	times:=make([]int64,len(timeStrs))
	for idx,val:=range timeStrs{
		time,_:=strconv.ParseInt(string(val),10,64)
		times[idx]=int64(time)
	}

	distStrs:=regex.FindAll([]byte(strings.Split(input[1], ":")[1]),-1)
	dists:=make([]int64,len(distStrs))
	for idx,val:=range distStrs{
		dist,_:=strconv.ParseInt(string(val),10,64)
		dists[idx]=int64(dist)
	}

	var ans int64=1
	for i:=0;i<len(times);i++{
		ans*=getNumPossible(times[i],dists[i])
	}
	fmt.Println(ans)
}

func getNumPossible(time int64,dist int64) int64 {
	var i int64=1
	for ;i<time;i++{
		distTraveled:=(time-i)*i
		if distTraveled>dist{
			break
		}
	}

	firstPos:=i
	lastPos:=time-i

	return lastPos-firstPos+1
}