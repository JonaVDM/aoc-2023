package main

import (
	"flag"
	"fmt"
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

	runners := []Runner{}

	for _, runner := range runners {
		if *onlyDay > 0 && runner.Day == *onlyDay {
			file := runner.File
			if *replacedInput != "" {
				file = *replacedInput
			}
			out := runner.Function(file)
			printOutput(runner.Day, out)
		}

		if *onlyDay == -1 {
			out := runner.Function(runner.File)
			printOutput(runner.Day, out)
		}
	}
}

func printOutput(day int, out [2]interface{}) {
	fmt.Printf("\n--- Day %d ---\nPart One: %v\nPart Two: %v\n", day, out[0], out[1])
}
