package main

import (
	"bufio"
	"container/list"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

var Direction = map[byte]int{
	'L': 0,
	'R': 1,
}

type Config struct {
	InputFile string
}

var cfg Config

func init() {
	flag.StringVar(&cfg.InputFile, "in", "", "input filename")
}

func ParseInput(line *string) (string, []string) {
	re := regexp.MustCompile(`\w+`)
	matches := re.FindAllString(*line, -1)
	return matches[0], []string{matches[1], matches[2]}
}

func DetermineSteps(start string, network *map[string][]string, instr *list.List) int {
	actual := start
	steps := 0
	for {
		for e := instr.Front(); e != nil; e = e.Next() {
			steps++
			actual = (*network)[actual][Direction[e.Value.(byte)]]
			if actual[len(actual)-1] == 'Z' {
				return steps
			}
		}
	}
}

// greatest common divisor (GCD)
// source: https://github.com/TheAlgorithms/Go/blob/master/math/gcd/gcditerative.go
func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
// source: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func Lcm(a, b int, integers ...int) int {
	result := a * b / Gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = Lcm(result, integers[i])
	}

	return result
}

func DetermineStepsPart2(network *map[string][]string, instr *list.List) int {
	var actual []string
	var steps []int

	// determine all nodes with ending 'A'
	for k := range *network {
		if k[2] == 'A' {
			actual = append(actual, k)
		}
	}

	// determine steps for every node
	for i := range actual {
		steps = append(steps, DetermineSteps(actual[i], network, instr))
	}

	// return steps
	if len(steps) > 2 {
		return Lcm(steps[0], steps[1], steps[2:]...)
	} else if len(steps) > 1 {
		return Lcm(steps[0], steps[1])
	} else {
		return steps[0]
	}
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

		instr := list.New() // instructions
		network := make(map[string][]string)

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()

			if len(line) == 0 { // skip blank lines
				continue
			} else if len(line) > 4 && line[4] == '=' {
				k, v := ParseInput(&line)
				network[k] = v
			} else { // instructions
				for i := range line {
					instr.PushBack(line[i])
				}
			}
		}

		fmt.Println("(part 1) steps from 'AAA' to 'ZZZ':", DetermineSteps("AAA", &network, instr))
		fmt.Println("(part 2) steps from 'xxA' to 'xxZ':", DetermineStepsPart2(&network, instr))

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
