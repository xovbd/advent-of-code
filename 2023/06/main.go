package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Config struct {
	InputFile string
}

var cfg Config

func ExtractValue(line *string) uint64 {
	s := strings.Split(*line, ":")[1]
	s = strings.ReplaceAll(s, " ", "")
	v, _ := strconv.ParseUint(s, 10, 64)
	return v
}

func ExtractValues(line *string) []uint64 {
	var vs []uint64
	var re = regexp.MustCompile(`\d+`)
	m := re.FindAllString(*line, -1)
	for i := range m {
		v, _ := strconv.Atoi(m[i])
		vs = append(vs, uint64(v))
	}
	return vs
}

func DetermineOpportunities(t, d *[]uint64) []uint64 {
	opps := make([]uint64, len(*t)) // number of opportunities per race
	for i := range *t {             // iterate over every racetime
		for ms := uint64(1); ms < (*t)[i]; ms++ { // iterate from 1ms to racetime-1ms
			if (ms * ((*t)[i] - ms)) > (*d)[i] {
				opps[i]++
			}
		}
	}
	return opps
}

func SolvePart1(opps *[]uint64) uint64 {
	var result uint64 = 1
	for i := range *opps {
		result *= (*opps)[i]
	}
	return result
}

func SolvePart2(t, d *uint64) uint64 {
	return DetermineOpportunities(&[]uint64{*t}, &[]uint64{*d})[0]
}

func init() {
	flag.StringVar(&cfg.InputFile, "in", "", "input filename")
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

		var time uint64
		var times []uint64
		var distance uint64
		var distances []uint64
		var opps []uint64

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()

			if line[0] == 'T' { // determine times
				times = ExtractValues(&line) // part1
				time = ExtractValue(&line)   // part2
			} else if line[0] == 'D' { // determine disctances
				distances = ExtractValues(&line) // part1
				distance = ExtractValue(&line)   // part2
			}
		}

		if len(times) != len(distances) {
			log.Fatalf("invalid input file: number times<%d> number distances<%d>\n",
				times, distances)
		} else {
			opps = DetermineOpportunities(&times, &distances)
		}

		fmt.Println("(part 1) multiplied opportunities:",
			SolvePart1(&opps))
		fmt.Println("(part 2)   possible opportunities:",
			SolvePart2(&time, &distance))

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
