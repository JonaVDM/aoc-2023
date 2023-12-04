package day04_test

import (
	"testing"

	"github.com/jonavdm/aoc-2023/day04"
	_ "github.com/jonavdm/aoc-2023/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{26426, 6227972}, day04.Run("day04"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day04.Run("day04")
	}
}
