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

type Card struct {
	Winning []int
	Numbers []int
	Score   []int
}

func getRealScore(input []string) int {

	cards := make(map[int]int)
	queue := make([]int, len(input))

	for i, line := range input {
		parts := strings.Split(line, ": ")
		card := make(map[int]int)
		score := 0

		nums := strings.Split(parts[1], " | ")

		idParts := strings.Split(parts[0], " ")
		id, _ := strconv.Atoi(idParts[len(idParts)-1])

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
			} else if _, ok := card[n]; ok {
				score++
			}
		}

		cards[id] = score
		queue[i] = id
	}

	total := 0
	for len(queue) > 0 {
		id := queue[0]
		queue = queue[1:]
		total++

		for i := id + 1; i <= id+cards[id] && i <= len(input); i++ {
			queue = append(queue, i)
		}
	}

	return total
}
