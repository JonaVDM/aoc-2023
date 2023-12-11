package day11

import (
	"github.com/jonavdm/aoc-2023/utils"
)

type Solver struct {
	Items     [][2]int
	EmptyRows []bool
	EmptyCols []bool
}

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	solver := parse(data)

	return [2]interface{}{
		solver.GetDistance(1),
		solver.GetDistance(999999),
	}
}

func parse(data []string) Solver {
	emptyRows := make([]bool, len(data[1]))
	emptyCols := make([]bool, len(data[1]))
	items := make([][2]int, 0)

	for i, row := range data {
		emptyRows[i] = true
		for j, col := range row {
			if col == '#' {
				items = append(items, [2]int{j, i})
				emptyRows[i] = false
			}
		}
	}

	for j := range data[0] {
		emptyCols[j] = true

		for i := range data {
			if data[i][j] == '#' {
				emptyCols[j] = false
			}
		}
	}

	return Solver{
		items, emptyRows, emptyCols,
	}
}

func (s *Solver) GetPairs() [][2]int {
	pairs := make([][2]int, 0)
	mappie := make(map[int]map[int]bool)

	for i := range s.Items {
		mappie[i] = make(map[int]bool)
		for j := range s.Items {
			if i == j {
				continue
			}

			if m, ok := mappie[j]; ok && m[i] {
				continue
			}

			pairs = append(pairs, [2]int{i, j})
			mappie[i][j] = true
		}
	}

	return pairs
}

func (s *Solver) GetDistance(expansion int) int {
	pairs := s.GetPairs()
	s.GetPairs()
	var total int

	for _, pair := range pairs {
		one := s.Items[pair[0]]
		two := s.Items[pair[1]]

		// up and down
		for i := min(one[1], two[1]); i < max(one[1], two[1]); i++ {
			total++
			if s.EmptyRows[i] {
				total += expansion
			}
		}

		// left and right
		for i := min(one[0], two[0]); i < max(one[0], two[0]); i++ {
			total++
			if s.EmptyCols[i] {
				total += expansion
			}
		}
	}

	return total
}
