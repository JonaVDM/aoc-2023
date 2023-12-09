package day05

import (
	"math"
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2023/utils"
)

type Solver struct {
	// dest, source start, source end
	Locations map[string][][3]int
	Seeds     []int
}

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)
	solver := parse(data)

	smallestA := math.MaxInt
	for _, seed := range solver.Seeds {
		soil := solver.GetNext(seed, "soil")
		fertilizer := solver.GetNext(soil, "fertilizer")
		water := solver.GetNext(fertilizer, "water")
		light := solver.GetNext(water, "light")
		temperature := solver.GetNext(light, "temperature")
		humidity := solver.GetNext(temperature, "humidity")
		location := solver.GetNext(humidity, "location")

		if location < smallestA {
			smallestA = location
		}
	}

	locations := []string{
		"soil", "fertilizer", "water", "light", "temperature", "humidity", "location",
	}

	smallestB := math.MaxInt
	for i := 0; i < len(solver.Seeds); i += 2 {
		if x := solver.FindRange([2]int{solver.Seeds[i], solver.Seeds[i] + solver.Seeds[i+1] - 1}, locations); x < smallestB {
			smallestB = x
		}
	}

	return [2]interface{}{
		smallestA,
		smallestB,
	}
}

func parse(input []string) Solver {
	solver := Solver{
		Locations: make(map[string][][3]int),
		Seeds:     make([]int, 0),
	}

	for _, seed := range strings.Split(input[0], " ")[1:] {
		n, _ := strconv.Atoi(seed)
		solver.Seeds = append(solver.Seeds, n)
	}

	next := true
	var current string
	for _, line := range input[2:] {
		if line == "" {
			next = true
			continue
		}

		if next {
			current = strings.Split(strings.Split(line, " ")[0], "-")[2]
			next = false
			solver.Locations[current] = make([][3]int, 0)
			continue
		}

		values := strings.Split(line, " ")
		dest, _ := strconv.Atoi(values[0])
		source, _ := strconv.Atoi(values[1])
		ran, _ := strconv.Atoi(values[2])

		solver.Locations[current] = append(solver.Locations[current], [3]int{dest, source, source + ran})
	}

	return solver
}

func (s *Solver) GetNext(source int, location string) int {
	for _, val := range s.Locations[location] {
		if source < val[1] || source > val[2] {
			continue
		}

		offset := source - val[1]
		return val[0] + offset
	}

	return source
}

func (s *Solver) GetNextRange(source [2]int, location string) [][2]int {
	queue := make([][2]int, 1)
	queue[0] = source

	out := make([][2]int, 0)

	for len(queue) > 0 {
		sor := queue[0]
		queue = queue[1:]
		br := false

		for _, m := range s.Locations[location] {
			// source is fully in map
			if sor[0] >= m[1] && sor[1] <= m[2] {
				between := sor[1] - sor[0]
				offset := sor[0] - m[1]
				out = append(out, [2]int{m[0] + offset, m[0] + offset + between})
				br = true
				break
			}

			// map is fully in source
			if m[1] >= sor[0] && m[2] <= sor[1] {
				out = append(out, [2]int{m[0], m[0] + m[2] - m[1] - 1})
				queue = append(queue, [2]int{sor[0], m[1] - 1})
				queue = append(queue, [2]int{m[2], sor[1]})
				br = true
				break
			}
		}

		if !br {
			out = append(out, sor)
		}
	}

	return out
}

func (s *Solver) FindRange(source [2]int, locations []string) int {
	if next := s.GetNextRange(source, locations[0]); len(locations) == 1 {
		smallest := next[0][0]
		for _, n := range next[1:] {
			if smallest > n[0] {
				smallest = n[0]
			}
		}
		return smallest
	} else {
		smallest := math.MaxInt
		for _, n := range next {
			num := s.FindRange(n, locations[1:])
			if num < smallest {
				smallest = num
			}
		}
		return smallest
	}
}
