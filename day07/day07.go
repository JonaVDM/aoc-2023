package day07

import (
	"sort"
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2023/utils"
)

type Hand struct {
	Cards     [5]int
	Bet       int
	Score     int
	JackScore int
}

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	hands := parse(data)
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Score == hands[j].Score {
			for k := 0; k < 5; k++ {
				if hands[i].Cards[k] == hands[j].Cards[k] {
					continue
				}

				return hands[i].Cards[k] < hands[j].Cards[k]
			}
		}

		return hands[i].Score < hands[j].Score
	})

	a := 0
	for i, card := range hands {
		a += card.Bet * (i + 1)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].JackScore == hands[j].JackScore {
			for k := 0; k < 5; k++ {
				iCard := hands[i].Cards[k]
				if iCard == 11 {
					iCard = 1
				}

				jCard := hands[j].Cards[k]
				if jCard == 11 {
					jCard = 1
				}

				if iCard == jCard {
					continue
				}

				return iCard < jCard
			}
		}

		return hands[i].JackScore < hands[j].JackScore
	})
	b := 0
	for i, card := range hands {
		b += card.Bet * (i + 1)
	}

	return [2]interface{}{
		a,
		b,
	}
}

func parse(data []string) []Hand {
	charToVal := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}

	hands := make([]Hand, len(data))
	for i, line := range data {
		spl := strings.Split(line, " ")
		bet, _ := strconv.Atoi(spl[1])

		hands[i] = Hand{
			Cards: [5]int{},
			Bet:   bet,
		}

		for j, char := range spl[0] {
			hands[i].Cards[j] = charToVal[char]
		}

		hands[i].Score = CalculateScore(hands[i].Cards, false)
		hands[i].JackScore = CalculateScore(hands[i].Cards, true)
	}
	return hands
}

func CalculateScore(cards [5]int, jack bool) int {
	// five of a kind (6)
	// four of a kind (5)
	// full house (4)
	// three of a kind (3)
	// two pairs (2)
	// pair (1)
	// high (0)

	largest := 0
	largestKey := 0
	count := make(map[int]int)

	for _, c := range cards {
		count[c]++
		if count[c] > largest && c != 11 {
			largestKey = c
			largest = count[c]
		}
	}

	fives := 0
	fours := 0
	pairs := 0
	trip := 0

	if jack && len(count) > 1 {
		count[largestKey] += count[11]
		count[11] = 0
	}

	for _, val := range count {
		if val >= 5 {
			fives++
		}

		if val == 4 {
			fours++
		}

		if val == 3 {
			trip++
		}

		if val == 2 {
			pairs++
		}
	}

	if fives == 1 {
		return 6
	} else if fours == 1 {
		return 5
	} else if pairs == 1 && trip == 1 {
		return 4
	} else if trip == 1 {
		return 3
	} else if pairs == 2 {
		return 2
	} else if pairs == 1 {
		return 1
	}

	return 0
}
