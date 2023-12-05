package day05

import (
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2023/utils"
)

type Location struct {
	Name        string
	Parent      string
	Destination []int
	Source      []int
	Range       []int
}

type Solver struct {
	Locations map[string]*Location
	Seeds     []int
}

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)
	solver := parse(data)

	smallestA := 10000000000000
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

	smallestB := 10000000000000
	for i := 0; i < len(solver.Seeds); i += 2 {
		for seed := solver.Seeds[i]; seed < solver.Seeds[i]+solver.Seeds[i+1]; seed++ {
			soil := solver.GetNext(seed, "soil")
			fertilizer := solver.GetNext(soil, "fertilizer")
			water := solver.GetNext(fertilizer, "water")
			light := solver.GetNext(water, "light")
			temperature := solver.GetNext(light, "temperature")
			humidity := solver.GetNext(temperature, "humidity")
			location := solver.GetNext(humidity, "location")

			if location < smallestB {
				smallestB = location
			}
		}
	}

	return [2]interface{}{
		smallestA,
		smallestB,
	}
}

func parse(input []string) Solver {
	solver := Solver{
		Locations: make(map[string]*Location),
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
			thing := Location{
				Name:        strings.Split(strings.Split(line, " ")[0], "-")[2],
				Parent:      strings.Split(strings.Split(line, " ")[0], "-")[0],
				Destination: make([]int, 0),
				Source:      make([]int, 0),
				Range:       make([]int, 0),
			}
			solver.Locations[current] = &thing
			next = false
			continue
		}

		values := strings.Split(line, " ")
		dest, _ := strconv.Atoi(values[0])
		source, _ := strconv.Atoi(values[1])
		ran, _ := strconv.Atoi(values[2])

		solver.Locations[current].Destination = append(solver.Locations[current].Destination, dest)
		solver.Locations[current].Source = append(solver.Locations[current].Source, source)
		solver.Locations[current].Range = append(solver.Locations[current].Range, ran)
	}

	return solver
}

func (s *Solver) GetNext(source int, location string) int {
	for i, val := range s.Locations[location].Source {
		if source < val || source > val+s.Locations[location].Range[i] {
			continue
		}

		offset := source - val
		return s.Locations[location].Destination[i] + offset
	}

	return source
}
