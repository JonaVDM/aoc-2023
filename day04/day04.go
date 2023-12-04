package day04

import (
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2023/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	return [2]interface{}{
		getScore(data),
		getRealScore(data),
	}
}

func getScore(input []string) int {
	total := 0

	for _, line := range input {
		parts := strings.Split(line, ": ")
		card := make(map[int]int)
		score := 0

		nums := strings.Split(parts[1], " | ")

		for _, num := range strings.Split(nums[1], " ") {
			if n, err := strconv.Atoi(num); err != nil {
				continue
			} else {
				card[n] += 1
			}
		}

		for _, num := range strings.Split(nums[0], " ") {
			if n, err := strconv.Atoi(num); err != nil {
				continue
			} else if _, ok := card[n]; !ok {
				continue
			} else if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}

		total += score
	}

	return total
}

func getRealScore(input []string) int {
	cardsAmount := make(map[int]int)
	total := 0

	for i, line := range input {
		parts := strings.Split(line, ": ")
		card := make(map[int]int)
		score := 0

		nums := strings.Split(parts[1], " | ")

		for _, num := range strings.Split(nums[1], " ") {
			if n, err := strconv.Atoi(num); err == nil {
				card[n] = 1
			}
		}

		for _, num := range strings.Split(nums[0], " ") {
			if n, err := strconv.Atoi(num); err == nil {
				continue
			} else if _, ok := card[n]; ok {
				score++
			}
		}

		cardsAmount[i] += 1
		total += cardsAmount[i]
		for j := i + 1; j <= i+score && j < len(input); j++ {
			cardsAmount[j] += cardsAmount[i]
		}
	}

	return total
}
