package day03

import (
	"fmt"
	"strconv"

	"github.com/jonavdm/aoc-2023/utils"
)

type Solver struct {
	Input       []string
	GearCounter map[string][]int
	Current     int
}

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)
	solver := Solver{
		Input:       data,
		GearCounter: make(map[string][]int),
	}

	return [2]interface{}{
		solver.findNumbers(),
		solver.findGears(),
	}
}

func (s *Solver) findNumbers() int {
	sum := 0

	for y, row := range s.Input {
		num := make([]int, 0)
		for x, col := range row {
			if col >= '0' && col <= '9' {
				num = append(num, x)
			} else {
				s.Current = s.convert(y, num)
				if len(num) > 0 && s.isAdjanced(y, num) {
					sum += s.convert(y, num)
					num = make([]int, 0)
				} else if len(num) > 0 {
					num = make([]int, 0)
				}
			}
		}

		s.Current = s.convert(y, num)
		if len(num) > 0 && s.isAdjanced(y, num) {
			sum += s.convert(y, num)
		}
	}

	return sum
}

func (s *Solver) convert(row int, cols []int) int {
	str := ""
	for _, n := range cols {
		str += string(s.Input[row][n])
	}
	n, _ := strconv.Atoi(str)
	return n
}

func (s *Solver) isAdjanced(row int, cols []int) bool {
	for _, col := range cols {
		adj := utils.GetAdjacend(col, row, len(s.Input), len(s.Input[0]))

		for _, point := range adj {
			if s.isSymbol(s.Input[point.Y][point.X], point.X, point.Y) {
				return true
			}
		}
	}

	return false
}

func (s *Solver) isSymbol(char byte, x, y int) bool {
	if char == '*' {
		str := fmt.Sprintf("%d,%d", x, y)
		if _, ok := s.GearCounter[str]; !ok {
			s.GearCounter[str] = []int{s.Current}
		} else {
			s.GearCounter[str] = append(s.GearCounter[str], s.Current)
		}
	}
	return char != '.' && (char < '0' || char > '9')
}

func (s *Solver) findGears() int {
	sum := 0
	for _, gear := range s.GearCounter {
		if len(gear) < 2 {
			continue
		}

		total := 1
		for _, thing := range gear {
			total *= thing
		}
		sum += total
	}

	return sum
}
