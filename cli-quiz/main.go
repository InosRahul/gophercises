package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilenanme := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilenanme)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFilenanme))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided csv file")
	}

	problems := parseLines(lines)
	correct := 0
	for i, p := range problems{
		fmt.Printf("Problem #%d: %s", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer{
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(lines))
}

func parseLines(lines [][]string) [] problem {
	ret := make([] problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answer string
}


func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}