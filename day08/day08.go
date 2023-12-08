package day08

import (
	"strings"

	"github.com/jonavdm/aoc-2023/utils"
)

type Solver struct {
	Data     map[string][2]string
	Sequence string
}

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	solver := parse(data)

	return [2]interface{}{
		solver.GetSteps(),
		solver.GetAllSteps(),
	}
}

func parse(data []string) Solver {
	solver := Solver{
		Data:     make(map[string][2]string),
		Sequence: data[0],
	}

	for _, line := range data[2:] {
		spl := strings.Split(line, " ")

		solver.Data[spl[0]] = [2]string{
			spl[2][1:4],
			spl[3][:3],
		}
	}

	return solver
}

func (s *Solver) GetSteps() int {
	position := "AAA"
	target := "ZZZ"
	steps := 0

	for {
		for _, instruction := range s.Sequence {
			steps++

			if instruction == 'R' {
				position = s.Data[position][1]
			} else {
				position = s.Data[position][0]
			}

			if position == target {
				return steps
			}
		}
	}
}

func (s *Solver) FindZ(starting string) int {
	position := starting
	steps := 0

	for {
		for _, instruction := range s.Sequence {
			steps++

			if instruction == 'R' {
				position = s.Data[position][1]
			} else {
				position = s.Data[position][0]
			}

			if position[2] == 'Z' {
				return steps
			}
		}
	}

}

func (s *Solver) GetAllSteps() int {
	postions := make([]int, 0)

	for key := range s.Data {
		if key[2] == 'A' {
			postions = append(postions, s.FindZ(key))
		}
	}

	return utils.LeastCommonDenominator(postions)
}
