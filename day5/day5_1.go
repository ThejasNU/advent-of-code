package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day5_1() {
	input := helpers.ReadInput("./day5/day5_input.txt")
	seedsStrs := strings.Split(input[0][7:], " ")
	seeds := make([]int64, len(seedsStrs))
	for idx, seedStr := range seedsStrs {
		seed, _ := strconv.ParseInt(seedStr, 10, 64)
		seeds[idx] = seed
	}

	seedSoilMap := getMap(input[2:35])
	soilFertilizerMap := getMap(input[36:60])
	fetilizerWaterMap := getMap(input[61:83])
	waterLightMap := getMap(input[84:103])
	lightTempMap := getMap(input[104:115])
	tempHumidMap := getMap(input[116:125])
	humidLocatMap := getMap(input[126:])

	soils := getDestinations(seeds, seedSoilMap)
	fertlizers := getDestinations(soils, soilFertilizerMap)
	waters := getDestinations(fertlizers, fetilizerWaterMap)
	lights := getDestinations(waters, waterLightMap)
	temps := getDestinations(lights, lightTempMap)
	humids := getDestinations(temps, tempHumidMap)
	locations := getDestinations(humids, humidLocatMap)

	nearestLocation := slices.Min(locations)

	fmt.Println(nearestLocation)
}

func getMap(ranges []string) map[int64][]int64 {
	//key:sourceStart, val:{destinationStart,length}
	m := make(map[int64][]int64)

	for _, s := range ranges {
		res := strings.Split(s, " ")
		sourceStart, _ := strconv.ParseInt(res[1], 10, 64)
		destinationStart, _ := strconv.ParseInt(res[0], 10, 64)
		length, _ := strconv.ParseInt(res[2], 10, 64)
		m[sourceStart] = []int64{destinationStart, length}
	}

	return m
}

func getDestinations(sources []int64, m map[int64][]int64) []int64 {
	destinations := make([]int64, len(sources))
	for idx, source := range sources {
		done := false
		for key, val := range m {
			if source >= key && source < key+val[1] {
				destinations[idx] = val[0] + source - key
				done = true
				break
			}
		}

		if !done {
			destinations[idx] = source
		}
	}
	return destinations
}
