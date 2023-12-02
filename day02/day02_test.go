package day02_test

import (
	"testing"

	"github.com/jonavdm/aoc-2023/day02"
	_ "github.com/jonavdm/aoc-2023/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{2237, 66681}, day02.Run("day02"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day02.Run("day02")
	}
}
