package day05_test

import (
	"testing"

	"github.com/jonavdm/aoc-2023/day05"
	_ "github.com/jonavdm/aoc-2023/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{57075758, 31161857}, day05.Run("day05"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day05.Run("day05")
	}
}

func TestGetNextRange(t *testing.T) {
	solver := day05.Solver{
		Locations: map[string][][3]int{
			"location": {
				[3]int{520, 100, 150},
				[3]int{600, 200, 250},
			},
		},
	}

	tests := []struct {
		source [2]int
		out    [][2]int
		label  string
	}{
		{[2]int{120, 130}, [][2]int{{540, 550}}, "source fully inside map"},
		{[2]int{80, 160}, [][2]int{{80, 99}, {520, 569}, {150, 160}}, "map fully inside source"},
		{[2]int{80, 260}, [][2]int{{80, 99}, {150, 199}, {250, 260}, {520, 569}, {600, 649}}, "two maps fully inside source"},
	}

	for _, test := range tests {
		assert.ElementsMatch(t, test.out, solver.GetNextRange(test.source, "location"), test.label)
	}
}
