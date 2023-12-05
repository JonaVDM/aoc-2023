package day05_test

import (
	"testing"

	"github.com/jonavdm/aoc-2023/day05"
	_ "github.com/jonavdm/aoc-2023/testing"
	// "github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	// assert.Equal(t, [2]interface{}{57075758, 31161857}, day05.Run("day05"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day05.Run("day05")
	}
}
