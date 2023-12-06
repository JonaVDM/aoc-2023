package day05

import (
	"log"
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
	channel := make(chan int)
	counter := 0

	for i := 0; i < len(solver.Seeds); i += 2 {
		go solver.FindValue(solver.Seeds[i], solver.Seeds[i+1], channel)
		counter++
	}

	log.Println("Total routines:", counter)

	for counter > 0 {
		val := <-channel
		counter--
		log.Println("Done with one!", counter, "to go")
		if val < smallestB {
			smallestB = val
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

func (s *Solver) FindValue(start, rang int, c chan int) {
	lowest := 10000000000000000
	for seed := start; seed < start+rang; seed++ {
		soil := s.GetNext(seed, "soil")
		fertilizer := s.GetNext(soil, "fertilizer")
		water := s.GetNext(fertilizer, "water")
		light := s.GetNext(water, "light")
		temperature := s.GetNext(light, "temperature")
		humidity := s.GetNext(temperature, "humidity")
		location := s.GetNext(humidity, "location")
		if location < lowest {
			lowest = location
		}
	}

	c <- lowest
}
