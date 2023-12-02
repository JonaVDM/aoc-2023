package day02

import (
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2023/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	a := 0
	b := 0
	for _, row := range data {
		if id, ok := solve(row); ok {
			a += id
		}

		b += getScore(row)
	}

	return [2]interface{}{
		a,
		b,
	}
}

func solve(row string) (int, bool) {
	maxs := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	game := strings.Split(row, ":")
	id, _ := strconv.Atoi(strings.Split(game[0], " ")[1])

	for _, set := range strings.Split(game[1], ";") {
		for _, ball := range strings.Split(set, ", ") {
			out := strings.Split(strings.TrimSpace(ball), " ")
			amount, _ := strconv.Atoi(out[0])

			if amount > maxs[out[1]] {
				return id, false
			}
		}
	}

	return id, true
}

func getScore(row string) int {
	game := strings.Split(row, ":")

	maxs := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, set := range strings.Split(game[1], ";") {
		for _, ball := range strings.Split(set, ", ") {
			out := strings.Split(strings.TrimSpace(ball), " ")
			amount, _ := strconv.Atoi(out[0])

			if amount > maxs[out[1]] {
				maxs[out[1]] = amount
			}
		}
	}

	return maxs["red"] * maxs["green"] * maxs["blue"]
}
