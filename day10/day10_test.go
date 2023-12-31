package day10_test

import (
	"testing"

	"github.com/jonavdm/aoc-2023/day10"
	_ "github.com/jonavdm/aoc-2023/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{6956, 0}, day10.Run("day10"))

	assert.Less(t, 681, day10.Run("day10")[1])
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day10.Run("day10")
	}
}

