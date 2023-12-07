package day07_test

import (
	"testing"

	"github.com/jonavdm/aoc-2023/day07"
	_ "github.com/jonavdm/aoc-2023/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{246912307, 0}, day07.Run("day07"))
}

func TestCalculateScore(t *testing.T) {
	tests := []struct {
		cards [5]int
		jack  bool
		out   int
		label string
	}{
		// without any jacks
		{[5]int{5, 5, 5, 5, 5}, false, 6, "five of a kind"},
		{[5]int{5, 5, 5, 5, 3}, false, 5, "four of a kind"},
		{[5]int{5, 5, 5, 3, 3}, false, 4, "full house"},
		{[5]int{5, 5, 5, 2, 1}, false, 3, "three of a kind"},
		{[5]int{5, 5, 2, 2, 1}, false, 2, "two pairs"},
		{[5]int{2, 2, 1, 9, 3}, false, 1, "pair"},
		{[5]int{1, 2, 3, 4, 5}, false, 0, "high"},

		// with jacks (jack has value 11)
		{[5]int{5, 5, 5, 5, 5}, true, 6, "(joker) fives of a kind (0 jokers)"},
		{[5]int{5, 5, 5, 5, 11}, true, 6, "(joker) fives of a kind (1 joker)"},
		{[5]int{5, 5, 11, 11, 5}, true, 6, "(joker) fives of a kind (2 jokers)"},
		{[5]int{5, 11, 11, 5, 11}, true, 6, "(joker) fives of a kind (3 jokers)"},
		{[5]int{11, 11, 5, 11, 11}, true, 6, "(joker) fives of a kind (4 jokers)"},
		{[5]int{11, 11, 11, 11, 11}, true, 6, "(joker) fives of a kind (5 jokers)"},

		{[5]int{5, 5, 5, 5, 3}, true, 5, "(joker) four of a kind (0 jokers)"},
		{[5]int{5, 5, 5, 11, 3}, true, 5, "(joker) four of a kind (1 joker)"},
		{[5]int{5, 5, 11, 11, 3}, true, 5, "(joker) four of a kind (2 jokers)"},
		{[5]int{5, 11, 11, 11, 3}, true, 5, "(joker) four of a kind (3 jokers)"},

		{[5]int{5, 5, 5, 3, 3}, true, 4, "(joker) full house (0 jokers)"},
		{[5]int{5, 5, 11, 3, 3}, true, 4, "(joker) full house (1 joker)"},

		{[5]int{5, 5, 5, 3, 4}, true, 3, "(joker) three of a kind (0 jokers)"},
		{[5]int{5, 5, 11, 3, 4}, true, 3, "(joker) three of a kind (1 joker)"},
		{[5]int{5, 11, 11, 3, 4}, true, 3, "(joker) three of a kind (2 jokers)"},

		{[5]int{5, 5, 3, 3, 4}, true, 2, "(joker) two pairs (0 jokers)"},

		{[5]int{5, 5, 3, 9, 4}, true, 1, "(joker) pair (0 jokers)"},
		{[5]int{5, 11, 3, 9, 4}, true, 1, "(joker) pair (1 joker)"},

		{[5]int{5, 0, 3, 9, 4}, true, 0, "(joker) pair (0 jokers)"},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, day07.CalculateScore(test.cards, test.jack), test.label)
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day07.Run("day07")
	}
}

