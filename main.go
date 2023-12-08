package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jonavdm/aoc-2023/day01"
	"github.com/jonavdm/aoc-2023/day02"
	"github.com/jonavdm/aoc-2023/day03"
	"github.com/jonavdm/aoc-2023/day04"
	// "github.com/jonavdm/aoc-2023/day05"
	"github.com/jonavdm/aoc-2023/day06"
	"github.com/jonavdm/aoc-2023/day07"
	"github.com/jonavdm/aoc-2023/day08"
)

type Runner struct {
	Day      int
	Function func(string) [2]interface{}
	File     string
}

func main() {
	onlyDay := flag.Int("day", -1, "Specify the day")
	replacedInput := flag.String("file", "", "Run with a different input")
	flag.Parse()

	runners := []Runner{
		{1, day01.Run, "day01"},
		{2, day02.Run, "day02"},
		{3, day03.Run, "day03"},
		{4, day04.Run, "day04"},
		// {5, day05.Run, "day05"},
		{6, day06.Run, "day06"},
		{7, day07.Run, "day07"},
		{8, day08.Run, "day08"},
	}

	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                -- Advent Of Code 2023 --                 ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("┏━━━━━┯━━━━━━━━━━━━━━━━━━┯━━━━━━━━━━━━━━━━━━┯━━━━━━━━━━━━━━┓")
	fmt.Println("┃ Day │ Part One         │ Part Two         │ Time         ┃")
	fmt.Println("┠─────┼──────────────────┼──────────────────┼──────────────┨")

	for _, runner := range runners {
		start := time.Now()
		var out [2]interface{}

		if *onlyDay > 0 && runner.Day == *onlyDay {
			file := runner.File
			if *replacedInput != "" {
				file = *replacedInput
			}
			out = runner.Function(file)
		}

		if *onlyDay == -1 {
			out = runner.Function(runner.File)
		}

		duration := time.Now().Sub(start)

		fmt.Printf("┃ %2d  │ %16v │ %16v │ %12v ┃\n", runner.Day, out[0], out[1], duration)
	}
	fmt.Println("┗━━━━━┷━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━┛")
}
