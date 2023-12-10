package day10

import (
	"github.com/jonavdm/aoc-2023/utils"
)

type Solver struct {
	Data      []string
	Pipes     [][]int
	PipeCount int
}

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	solver := Solver{
		Data:  data,
		Pipes: make([][]int, len(data)),
	}
	for i, row := range data {
		solver.Pipes[i] = make([]int, len(row))
	}

	return [2]interface{}{
		solver.GetFurtest(),
		solver.GetEmpty(),
	}
}

func (s *Solver) FindStart() [2]int {
	for y, row := range s.Data {
		for x, col := range row {
			if col == 'S' {
				return [2]int{x, y}
			}
		}
	}
	return [2]int{-1, -1}
}

func (s *Solver) GetNext(current [2]int) [][2]int {
	out := make([][2]int, 0)
	if current[0] < 0 || current[1] < 0 || current[0] >= len(s.Data[1]) || current[1] >= len(s.Data) {
		return out
	}

	up := [2]int{current[0], current[1] - 1}
	down := [2]int{current[0], current[1] + 1}
	left := [2]int{current[0] - 1, current[1]}
	right := [2]int{current[0] + 1, current[1]}

	char := s.Data[current[1]][current[0]]

	switch char {
	case 'S':
		upCheck := up[1] >= 0 && up[1] < len(s.Data) && utils.RuneInString(rune(s.Data[up[1]][up[0]]), "7F|")
		downCheck := down[1] >= 0 && down[1] < len(s.Data) && utils.RuneInString(rune(s.Data[down[1]][down[0]]), "JL|")
		leftCheck := left[0] >= 0 && left[0] < len(s.Data) && utils.RuneInString(rune(s.Data[left[1]][left[0]]), "LF-")
		rightCheck := right[0] >= 0 && right[0] < len(s.Data) && utils.RuneInString(rune(s.Data[right[1]][right[0]]), "7J-")
		if upCheck {
			out = append(out, up)
		}
		if downCheck {
			out = append(out, down)
		}
		if leftCheck {
			out = append(out, left)

		}
		if rightCheck {
			out = append(out, right)
		}

	case '|':
		out = append(out, up)
		out = append(out, down)

	case '-':
		out = append(out, left)
		out = append(out, right)

	case 'L':
		out = append(out, up)
		out = append(out, right)

	case 'J':
		out = append(out, up)
		out = append(out, left)

	case '7':
		out = append(out, down)
		out = append(out, left)

	case 'F':
		out = append(out, down)
		out = append(out, right)
	}

	return out
}

func (s *Solver) GetFurtest() int {
	queue := make([][2]int, 1)
	queue[0] = s.FindStart()
	s.Pipes[queue[0][1]][queue[0][0]] = 1

	count := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		value := s.Pipes[current[1]][current[0]]
		count++

		for _, child := range s.GetNext(current) {
			if s.Pipes[child[1]][child[0]] == 0 {
				s.Pipes[child[1]][child[0]] = value + 1
				queue = append(queue, child)
			}
		}
	}

	s.PipeCount = count
	return count / 2
}

func (s *Solver) Checkpipes(x1, x2, y1, y2, xDir, yDir int) ([2]int, bool) {
	for {
		if y1 < 0 || y1 >= len(s.Data) || x1 < 0 || x1 >= len(s.Data[1]) {
			break
		}

		next1 := s.Pipes[y1][x1]
		next2 := s.Pipes[y2][x2]

		if utils.AbsInt(next1-next2) == 1 {
			break
		}

		if next1 <= 0 {
			return [2]int{x1, y1}, true
		}

		if next2 <= 0 {
			return [2]int{x2, y2}, true
		}

		x1 += xDir
		x2 += xDir
		y1 += yDir
		y2 += yDir
	}

	return [2]int{}, false
}

func (s *Solver) NextEmpty(x, y int) [][2]int {
	out := make([][2]int, 0)

	up := [2]int{x, y - 1}
	down := [2]int{x, y + 1}
	left := [2]int{x - 1, y}
	right := [2]int{x + 1, y}

	upEmpty := up[1] >= 0
	downEmpty := down[1] < len(s.Data)
	leftEmpty := left[0] >= 0
	rightEmpty := right[0] < len(s.Data)

	if upEmpty {
		out = append(out, up)
	}

	if downEmpty {
		out = append(out, down)
	}

	if leftEmpty {
		out = append(out, left)
	}

	if rightEmpty {
		out = append(out, right)
	}

	if leftEmpty && upEmpty {
		if data, ok := s.Checkpipes(x-1, x-1, y, y-1, -1, 0); ok {
			out = append(out, data)
		}

		if data, ok := s.Checkpipes(x-1, x, y-1, y-1, 0, -1); ok {
			out = append(out, data)
		}
	}

	if rightEmpty && upEmpty {
		if data, ok := s.Checkpipes(x+1, x, y-1, y-1, 0, -1); ok {
			out = append(out, data)
		}

		if data, ok := s.Checkpipes(x+1, x+1, y, y-1, 1, 0); ok {
			out = append(out, data)
		}
	}

	if leftEmpty && downEmpty {
		if data, ok := s.Checkpipes(x-1, x, y+1, y+1, 0, 1); ok {
			out = append(out, data)
		}

		if data, ok := s.Checkpipes(x-1, x-1, y+1, y, -1, 0); ok {
			out = append(out, data)
		}
	}

	if rightEmpty && downEmpty {
		if data, ok := s.Checkpipes(x+1, x+1, y, y+1, 1, 0); ok {
			out = append(out, data)
		}

		if data, ok := s.Checkpipes(x+1, x, y+1, y+1, 0, 1); ok {
			out = append(out, data)
		}
	}

	return out
}

func (s *Solver) GetEmpty() int {
	queue := make([][2]int, 0)
	for i := range s.Data[0] {
		queue = append(queue, [2]int{i, len(s.Data) - 1})
		queue = append(queue, [2]int{i, 0})
	}

	for i := range s.Data {
		queue = append(queue, [2]int{0, i})
		queue = append(queue, [2]int{len(s.Data[0]) - 1, i})
	}

	counter := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if s.Pipes[current[1]][current[0]] != 0 {
			continue
		}

		s.Pipes[current[1]][current[0]] = -1
		counter++

		for _, child := range s.NextEmpty(current[0], current[1]) {
			queue = append(queue, child)
		}
	}

	return len(s.Data)*len(s.Data[1]) - s.PipeCount - counter
}
