package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type Config struct {
	InputFile string
}

var cfg Config

func init() {
	flag.StringVar(&cfg.InputFile, "in", "", "input filename")
}

func ParseInput(line *string) []int {
	var seq []int
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(*line, -1)
	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		seq = append(seq, num)
	}
	return seq
}

func ExtrapolateSeq(seq []int) int {
	var diffSeq []int
	for i := 0; i < len(seq)-1; i++ {
		diffSeq = append(diffSeq, seq[i+1]-seq[i])
	}

	// check if all values in slice are 0.0
	if !slices.ContainsFunc(diffSeq, func(x int) bool { return x != 0 }) {
		return seq[len(seq)-1]
	}

	//fmt.Println(diffSeq)

	return seq[len(seq)-1] + ExtrapolateSeq(diffSeq)
}

func main() {

	flag.Parse()

	// check passed arguments
	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name] = true })
	if !flagset["in"] {
		fmt.Println("you have to pass all arguments:")
		flag.PrintDefaults()
	} else {

		file, err := os.Open(cfg.InputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		var sumPart1, sumPart2 int

		for scanner.Scan() {
			line := scanner.Text()

			seq := ParseInput(&line)
			//fmt.Println(seq)
			sumPart1 += ExtrapolateSeq(seq)
			slices.Reverse(seq)
			//fmt.Println(seq)
			sumPart2 += ExtrapolateSeq(seq)
		}

		fmt.Println("(part 1)", sumPart1)
		fmt.Println("(part 2)", sumPart2)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
