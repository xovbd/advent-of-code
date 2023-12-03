package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Config struct {
	InputFile string
}

var cfg Config

type Point struct {
	X, Y int
}

type Frame struct {
	P0 Point
	P1 Point
}

type Number struct {
	Value int
	Frame Frame
}

type Symbol struct {
	Char byte
	Y    int
	X    int
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

		// regex to find number
		var reNum = regexp.MustCompile(`\d+`)

		// regex to find symbol
		var reSym = regexp.MustCompile(`[^.\d]+`)

		var foundNumbers []Number
		var foundSymbols []Symbol

		scanner := bufio.NewScanner(file)

		col := 0
		for scanner.Scan() {
			col += 1
			line := scanner.Text()

			// search numbers
			matches := reNum.FindAllStringIndex(line, -1)
			for _, match := range matches {
				number, _ := strconv.Atoi(line[match[0]:match[1]])
				foundNumbers = append(foundNumbers,
					Number{
						Value: number,
						Frame: Frame{
							P0: Point{X: match[0], Y: col - 1},
							P1: Point{X: match[1] + 1, Y: col + 1},
						},
					},
				)
			}

			// search symbols
			matches = reSym.FindAllStringIndex(line, -1)
			for _, match := range matches {
				symbol := line[match[0]]
				foundSymbols = append(foundSymbols, Symbol{symbol, col, match[0] + 1})
			}
		}

		var sumValidNumbers int
		var sumGearRatios int

		// determine valid numbers
		for _, number := range foundNumbers {
			validNumber := false

			for _, symbol := range foundSymbols {
				// check if symbol is in frame
				if symbol.Y >= number.Frame.P0.Y &&
					symbol.Y <= number.Frame.P1.Y &&
					symbol.X >= number.Frame.P0.X &&
					symbol.X <= number.Frame.P1.X {
					validNumber = true
					break
				}
			}
			if validNumber {
				sumValidNumbers += number.Value
			}
		}

		// determine gear numbers
		for _, symbol := range foundSymbols {

			if symbol.Char == '*' {
				var gearNumbers []int

				for _, number := range foundNumbers {
					// check if symbol is in frame
					if symbol.Y >= number.Frame.P0.Y &&
						symbol.Y <= number.Frame.P1.Y &&
						symbol.X >= number.Frame.P0.X &&
						symbol.X <= number.Frame.P1.X {
						gearNumbers = append(gearNumbers, number.Value)
					}
				}
				if len(gearNumbers) > 1 {
					gearRatio := 0
					for _, value := range gearNumbers {
						if gearRatio == 0 {
							gearRatio = value
						} else {
							gearRatio *= value
						}
					}
					sumGearRatios += gearRatio
				}
			}
		}

		fmt.Println("(part 1) sum of valid numbers:", sumValidNumbers)
		fmt.Println("(part 2) sum of gear ratios  :", sumGearRatios)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
