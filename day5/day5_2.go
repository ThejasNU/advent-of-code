package day5

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day5_2() {
	input := helpers.ReadInput("./day5/day5_input.txt")

	soilSeedMap := getInvMap(input[2:35])
	fertlizerSoilMap := getInvMap(input[36:60])
	waterFertilizerMap := getInvMap(input[61:83])
	lightWaterMap := getInvMap(input[84:103])
	tempLightMap := getInvMap(input[104:115])
	humidTempMap := getInvMap(input[116:125])
	locatHumidMap := getInvMap(input[126:])

	seedsStrs := strings.Split(input[0][7:], " ")
	seedRanges := make(map[int64]int64)
	for i := 0; i < len(seedsStrs); i += 2 {
		seedStart, _ := strconv.ParseInt(seedsStrs[i], 10, 64)
		length, _ := strconv.ParseInt(seedsStrs[i+1], 10, 64)
		seedEnd := seedStart + length - 1
		seedRanges[seedStart] = seedEnd
	}

	var location int64 = 0
	for ; location < math.MaxInt64; location++ {
		humid := getDestination(location, locatHumidMap)
		temp := getDestination(humid, humidTempMap)
		light := getDestination(temp, tempLightMap)
		water := getDestination(light, lightWaterMap)
		fetilizer := getDestination(water, waterFertilizerMap)
		soil := getDestination(fetilizer, fertlizerSoilMap)
		seed := getDestination(soil, soilSeedMap)

		if isValid(seed, seedRanges) {
			break
		}
	}

}

func getInvMap(ranges []string) map[int64][]int64 {
	//key:sourceStart, val:{destinationStart,length}
	m := make(map[int64][]int64)

	for _, s := range ranges {
		res := strings.Split(s, " ")
		sourceStart, _ := strconv.ParseInt(res[1], 10, 64)
		destinationStart, _ := strconv.ParseInt(res[0], 10, 64)
		length, _ := strconv.ParseInt(res[2], 10, 64)
		m[destinationStart] = []int64{sourceStart, length}
	}

	return m
}

func getDestination(source int64, m map[int64][]int64) int64 {
	destination := source
	for key, val := range m {
		if source >= key && source < key+val[1] {
			destination = val[0] + source - key
			break
		}
	}
	return destination
}

func isValid(seed int64, seedRanges map[int64]int64) bool {
	for start, end := range seedRanges {
		if start > seed {
			return false
		}

		if seed >= start && seed <= end {
			fmt.Println(seed)
			return true
		}
	}

	return false
}
