package day09_test

import (
	"testing"

	"github.com/jonavdm/aoc-2023/day09"
	_ "github.com/jonavdm/aoc-2023/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{1743490457, 1053}, day09.Run("day09"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day09.Run("day09")
	}
}
