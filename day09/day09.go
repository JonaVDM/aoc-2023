package day09

import (
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2023/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	var next, prev int

	for _, line := range data {
		n, p := getNextNumber(line)
		next += n
		prev += p
	}

	return [2]interface{}{
		next,
		prev,
	}
}

func getNextNumber(line string) (int, int) {
	parts := strings.Split(line, " ")
	layers := make([][]int, 1)

	layers[0] = make([]int, len(parts))
	for i, part := range parts {
		layers[0][i], _ = strconv.Atoi(part)
	}

	i := 0
	for {
		layer, zero := findNextLayer(layers[i])
		layers = append(layers, layer)
		if zero {
			break
		}
		i++
	}

	next := 0
	prev := 0
	for j := i; j >= 0; j-- {
		next = layers[j][len(layers[j])-1] + next
		prev = layers[j][0] - prev
	}

	return next, prev
}

func findNextLayer(layer []int) ([]int, bool) {
	zeros := true
	next := make([]int, len(layer)-1)

	for i := range next {
		next[i] = layer[i+1] - layer[i]
		if zeros && next[i] != 0 {
			zeros = false
		}
	}

	return next, zeros
}
