package day01

import (
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2023/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	a := 0
	b := 0
	for _, l := range data {
		a += count(l)
		b += count(formatString(l))
	}

	return [2]interface{}{
		a,
		b,
	}
}

func count(line string) int {
	num := ""
	for i := 0; i < len(line); i++ {
		_, err := strconv.Atoi(string(line[i]))
		if err != nil {
			continue
		}

		num += string(line[i])
		break
	}

	for i := len(line) - 1; i >= 0; i-- {
		_, err := strconv.Atoi(string(line[i]))
		if err != nil {
			continue
		}

		num += string(line[i])
		break
	}

	total, _ := strconv.Atoi(num)
	return total
}

func formatString(input string) string {
	input = strings.ReplaceAll(input, "one", "on1e")
	input = strings.ReplaceAll(input, "two", "tw2o")
	input = strings.ReplaceAll(input, "three", "thr3ee")
	input = strings.ReplaceAll(input, "four", "fo4ur")
	input = strings.ReplaceAll(input, "five", "fi5ve")
	input = strings.ReplaceAll(input, "six", "si6x")
	input = strings.ReplaceAll(input, "seven", "sev7en")
	input = strings.ReplaceAll(input, "eight", "eig8ht")
	input = strings.ReplaceAll(input, "nine", "ni9ne")

	return input
}
