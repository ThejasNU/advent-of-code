package day7

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/ThejasNU/advent-of-code/helpers"
)

func Day7_1() {
	input := helpers.ReadInput("./day7/day7_input.txt")
	regex, _ := regexp.Compile("[0-9]+")
	cardsStrengthMap := getCardsStrengthMap()
	cardsTypes := make([][][]string, 7)

	for _, s := range input {
		cards := s[0:5]
		betBytes := regex.Find([]byte(s[6:]))
		bet := string(betBytes)

		idx := getIdxOfDeck(cards)

		cardsTypes[idx] = append(cardsTypes[idx], []string{cards, bet})
	}

	//sorting
	for idx, deckType := range cardsTypes {
		sort.SliceStable(deckType, func(i, j int) bool {
			k := 0
			string1 := deckType[i][0]
			string2 := deckType[j][0]
			for k < 5 && string1[k] == string2[k] {
				k++
			}

			if k >= 5 {
				return true
			} else {
				s1, _ := cardsStrengthMap[rune(string1[k])]
				s2, _ := cardsStrengthMap[rune(string2[k])]

				return s1 < s2
			}
		})
		cardsTypes[idx] = deckType
	}

	printAns(cardsTypes)
}

/*
high card:0
one pair:1
two pair:2
three of a kind:3
full house:4
four of a kind:5
five of a kind:6
*/
func getIdxOfDeck(cards string) int {
	m := make(map[rune]int)
	for _, card := range cards {
		if val, ok := m[card]; ok {
			m[card] = val + 1
		} else {
			m[card] = 1
		}
	}

	numDifCards := len(m)

	if numDifCards == 5 {
		return 0
	} else if numDifCards == 4 {
		return 1
	} else if numDifCards == 3 {
		numSingleCard := 0

		for _, val := range m {
			if val == 1 {
				numSingleCard++
			}
		}

		if numSingleCard == 1 {
			return 2
		} else {
			return 3
		}
	} else if numDifCards == 2 {
		numSingleCard := 0

		for _, val := range m {
			if val == 1 {
				numSingleCard++
			}
		}

		if numSingleCard == 0 {
			return 4
		} else {
			return 5
		}
	} else {
		return 6
	}
}

func printAns(cardTypes [][][]string) {
	typeIdx := 0
	cardDeckIdx := 0
	var rank int64 = 1
	var ans int64 = 0
	for typeIdx < 7 {
		for cardDeckIdx < len(cardTypes[typeIdx]) {
			bet, _ := strconv.ParseInt(cardTypes[typeIdx][cardDeckIdx][1], 10, 64)
			ans += rank * bet
			rank++
			cardDeckIdx++
		}
		typeIdx++
		cardDeckIdx = 0
	}

	fmt.Println(ans)
}

func getCardsStrengthMap() map[rune]int {
	m := make(map[rune]int)

	m['A'] = 13
	m['K'] = 12
	m['Q'] = 11
	m['J'] = 10
	m['T'] = 9
	m['9'] = 8
	m['8'] = 7
	m['7'] = 6
	m['6'] = 5
	m['5'] = 4
	m['4'] = 3
	m['3'] = 2
	m['2'] = 1

	return m
}
