package day06

import (
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2023/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	time := make([]int, 0)
	for _, part := range strings.Split(data[0], " ")[1:] {
		if part == "" {
			continue
		}
		num, _ := strconv.Atoi(part)
		time = append(time, num)
	}

	distance := make([]int, 0)
	for _, part := range strings.Split(data[1], " ")[1:] {
		if part == "" {
			continue
		}
		num, _ := strconv.Atoi(part)
		distance = append(distance, num)
	}

	a := 1
	for i := range distance {
		a *= findScore(time[i], distance[i])
	}

	longTime, _ := strconv.Atoi(strings.ReplaceAll(data[0][5:], " ", ""))
	longDistance, _ := strconv.Atoi(strings.ReplaceAll(data[1][9:], " ", ""))

	return [2]interface{}{
		a,
		findScore(longTime, longDistance),
	}
}

func findScore(time, score int) int {
	var start int

	for i := 0; i <= time; i++ {
		remaining := time - i
		if score < remaining*i {
			start = i
			break
		}
	}

	return time - start*2 + 1
}
