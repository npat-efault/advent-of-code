package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strconv"
)

type dayFunc func(input io.Reader)

var days = [...]dayFunc{
	Day01,
	Day02,
	Day03,
	Day04,
	Day05,
}

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)
	log.SetPrefix(path.Base(os.Args[0]) + ": ")

	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Printf("Usage is: %s <day #> [<input file>]\n",
			path.Base(os.Args[0]))
		os.Exit(1)
	}
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}
	if day < 1 || day > len(days) {
		log.Fatalf("Day %d not implemented", day)
	}
	input, infile := os.Stdin, "<stdin>"
	if len(os.Args) > 2 {
		input, err = os.Open(os.Args[2])
		if err != nil {
			log.Fatal(err.Error())
		}
		infile = os.Args[2]
	}
	log.Printf("Running day %d, with input from: %s", day, infile)
	log.SetPrefix(fmt.Sprintf("Day %02d: ", day))
	days[day-1](input)
}
