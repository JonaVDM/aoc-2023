package day01_test

import (
	"testing"

	"github.com/jonavdm/aoc-2023/day01"
	_ "github.com/jonavdm/aoc-2023/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{55017, 53539}, day01.Run("day01"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day01.Run("day01")
	}
}
